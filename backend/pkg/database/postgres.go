package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/shishir54234/NewsScraper/backend/pkg/utils"
	"gorm.io/gorm"
)
type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"dbName"`
	SSLMode  bool   `mapstructure:"sslMode"`
	Password string `mapstructure:"password"`
}
type PostgresDB struct {
	DB *sqlx.DB
	config *PostgresConfig
}	

func NewPostgresDB(dataSourceName string, config *PostgresConfig) (*PostgresDB, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * 60) // 5 minutes
	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresDB{DB: db, config: config}, nil
}

func (db *PostgresDB) Close() error {
	if db.DB != nil {
		return db.DB.Close()
	}
	return nil
}

func (db *PostgresDB) GetDB() *sqlx.DB {
	return db.DB
}

func (db* PostgresDB) Health() error{
	return db.DB.Ping()
}

func Migrate(gorm *gorm.DB, types ...interface{}) error {

	for _, t := range types {
		err := gorm.AutoMigrate(t)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *PostgresDB) Transaction(fn func(*sqlx.Tx)) (error) {
	tx, err := db.DB.Beginx()
	
	
	
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if r := recover(); r != nil {
			if err := tx.Rollback(); err != nil {
				fmt.Printf("failed to rollback transaction: %v\n", err)
			}
			panic(r) // re-throw panic after rollback
		}else if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("failed to rollback transaction: %v\n", rbErr)
			}
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				fmt.Printf("failed to commit transaction: %v\n", commitErr)
			}
		}
	}()
	fn(tx)
	return nil
}
func Paginate[T any](ctx context.Context,
	listQuery *utils.ListQuery, db *gorm.DB) (*utils.ListResult[T], error) {

	var items []T
	var totalRows int64
	db.Model(items).Count(&totalRows)

	// generate where query
	query := db.Offset(listQuery.GetOffset()).Limit(listQuery.GetLimit()).Order(listQuery.GetOrderBy())

	if listQuery.Filters != nil {
		for _, filter := range listQuery.Filters {
			column := filter.Field
			action := filter.Comparison
			value := filter.Value

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				query = query.Where(whereQuery, value)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				query = query.Where(whereQuery, "%"+value+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(value, ",")
				query = query.Where(whereQuery, queryArray)
				break

			}
		}
	}
	if err := query.Find(&items).Error; err != nil {
		return nil, errors.Wrap(err, "error in finding products.")
	}

	return utils.NewListResult[T](items, listQuery.GetSize(), listQuery.GetPage(), totalRows), nil

}