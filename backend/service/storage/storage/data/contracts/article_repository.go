package contracts

import (
	"context"

	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	utils "github.com/shishir54234/NewsScraper/backend/pkg/utils"
)


type ArticleRepository interface {
	GetAllArticles(ctx context.Context, listQuery *utils.ListQuery) (*utils.ListResult[*models.Article], error)
	CreateArticle(ctx context.Context, article *models.Article) (*models.Article, error)
	
}