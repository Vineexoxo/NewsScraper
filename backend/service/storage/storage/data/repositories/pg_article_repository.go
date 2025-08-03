package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	pgsql "github.com/shishir54234/NewsScraper/backend/pkg/database"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/utils"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/data/contracts"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PostgresArticleRepository struct {
	log  logger.ILogger
	cfg  *pgsql.PostgresConfig
	db   *pgxpool.Pool
	gorm *gorm.DB
}

func NewPostgresArticleRepository(log logger.ILogger, cfg *pgsql.PostgresConfig, gorm *gorm.DB) contracts.ArticleRepository {
	return &PostgresArticleRepository{log: log, cfg: cfg, gorm: gorm}
}

func (p *PostgresArticleRepository) GetAllArticles(ctx context.Context, listQuery *utils.ListQuery) (*utils.ListResult[*models.Article], error) {

	result, err := pgsql.Paginate[*models.Article](ctx, listQuery, p.gorm)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PostgresArticleRepository) SearchArticles(ctx context.Context, searchText string, listQuery *utils.ListQuery) (*utils.ListResult[*models.Article], error) {

	whereQuery := fmt.Sprintf("%s IN (?)", "Name")
	query := p.gorm.Where(whereQuery, searchText)

	result, err := pgsql.Paginate[*models.Article](ctx, listQuery, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func (p *PostgresArticleRepository) GetArticleById(ctx context.Context, uuid uuid.UUID) (*models.Article, error) {

	var article models.Article

	if err := p.gorm.First(&article, uuid).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't find the article with id %s into the database.", uuid))
	}

	return &article, nil
}

func (p *PostgresArticleRepository) CreateArticle(ctx context.Context, article *models.Article) (*models.Article, error) {

	if err := p.gorm.Create(&article).Error; err != nil {
		return nil, errors.Wrap(err, "error in the inserting article into the database.")
	}

	return article, nil
}

func (p *PostgresArticleRepository) UpdateArticle(ctx context.Context, updateArticle *models.Article) (*models.Article, error) {

	if err := p.gorm.Save(updateArticle).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error in updating article with id %s into the database."))
	}

	return updateArticle, nil
}

func (p *PostgresArticleRepository) DeleteArticleByID(ctx context.Context, uuid uuid.UUID) error {

	var article models.Article

	if err := p.gorm.First(&article, uuid).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("can't find the article with id %s into the database.", uuid))
	}

	if err := p.gorm.Delete(&article).Error; err != nil {
		return errors.Wrap(err, "error in the deleting article into the database.")
	}

	return nil
}
