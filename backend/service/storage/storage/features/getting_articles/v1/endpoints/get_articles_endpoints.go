package endpoints

import
(
	"context"
	"github.com/go-playground/validator"
	echo "github.com/labstack/echo/v4"
	mediatr "github.com/mehdihadeli/go-mediatr"
	logger "github.com/shishir54234/NewsScraper/backend/pkg"
	commandsv1 "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/commands"
	
	dtosv1 "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/dtos"
	"github.com/pkg/errors"
	"net/http"
)

func Maproute(ctx context.Context, validator *validator.Validate,log logger.ILogger, e *echo.Echo) {
	group:= e.Group("/api/v1/articles/get")
		group.GET(
			"/",
			getArticlesHandler(ctx), 
		)
}
	
	// getArticlesHandler is an example handler for the articles endpoint.
func getArticlesHandler(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error{
		request:= &dtosv1.RequestArticleDto{}
		if err:= c.Bind(request); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		result, err := mediatr.Send(ctx, commandsv1.NewGetArticlesQuery(request.URL))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"result": result,
		})


	}
}

