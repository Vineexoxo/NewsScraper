package contracts

import(
	"context"
	"github.com/shishir54234/NewsScraper/backend/pkg/utils"
)


type ArticleRepository interface {
	GetAllArticles(ctx context.Context, listQuery *utils.ListQuery) (*utils.ListResult[*models.Product], error)


}