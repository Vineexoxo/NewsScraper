package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

)

type PostgresDB struct {
	DB *sqlx.DB
}	

func NewPostgresDB(dataSourceName string) (*PostgresDB, error) {
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

	return &PostgresDB{DB: db}, nil
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




