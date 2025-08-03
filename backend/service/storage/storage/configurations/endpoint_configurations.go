package configurations

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	getting_products "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/endpoints"
)

func ConfigEndpoints(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {

	getting_products.Maproute(ctx, validator, log, echo)
	// creating_product.MapRoute(validator, log, echo, ctx)
	// deleting_product.MapRoute(validator, log, echo, ctx)
	// getting_product_by_id.MapRoute(validator, log, echo, ctx)
	// searching_product.MapRoute(validator, log, echo, ctx)
	// updating_product.MapRoute(validator, log, echo, ctx)
}
