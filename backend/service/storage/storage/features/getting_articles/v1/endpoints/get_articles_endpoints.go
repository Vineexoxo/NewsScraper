package endpoints

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
	echo "github.com/labstack/echo/v4"
	mediatr "github.com/mehdihadeli/go-mediatr"
	logger "github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/utils"
	commandsv1 "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/commands"

	dtosv1 "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/dtos"

	"net/http"
)

func Maproute(ctx context.Context, validator *validator.Validate,log logger.ILogger, e *echo.Echo) {
	group:= e.Group("/api/v1/articles/get")
		group.GET(
			"/",
			getArticlesHandler(ctx), 
		)
	for _, r := range e.Routes() {
		fmt.Printf("Method: %s, Path: %s\n", r.Method, r.Path)
	}
}
	
	// getArticlesHandler is an example handler for the articles endpoint.
func getArticlesHandler(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error{
		listQuery, err:=utils.GetListQueryFromCtx(c)
		
		if err!=nil{ return c.JSON(http.StatusBadRequest, err) }
		getArticlesCommand := commandsv1.NewGetArticles(listQuery)
		
		result, err := mediatr.Send[*commandsv1.GetArticles, *dtosv1.ResponseArticleDto](ctx, getArticlesCommand)
		if err!=nil{ return c.JSON(http.StatusInternalServerError, err) }
		return c.JSON(http.StatusOK, result)
	}
}

