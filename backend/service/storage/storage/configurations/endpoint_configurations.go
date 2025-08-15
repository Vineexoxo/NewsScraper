package configurations

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	getting_articles "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/endpoints"
	creating_article "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/creating_article/v1/endpoints"
	getting_articles_by_url "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/get_article_by_url/endpoints"
)

func ConfigEndpoints(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {

	getting_articles.Maproute(ctx, validator, log, echo)
	creating_article.MapRoute(validator, log, echo, ctx)
	getting_articles_by_url.MapRoute(ctx, validator, log, echo)
	
	// deleting_product.MapRoute(validator, log, echo, ctx)
	// getting_product_by_id.MapRoute(validator, log, echo, ctx)
	// searching_product.MapRoute(validator, log, echo, ctx)
	// updating_product.MapRoute(validator, log, echo, ctx)
}
