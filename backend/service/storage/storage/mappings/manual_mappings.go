package mappings

import (
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/dtos"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
)

func ProductToProductResponseDto(product *models.Article) *dtos.RequestArticleDto {
	return &dtos.RequestArticleDto{}
}
