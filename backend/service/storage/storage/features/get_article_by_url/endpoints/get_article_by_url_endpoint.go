package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	echo "github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	dtos "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/get_article_by_url/dtos"
)

func MapRoute(ctx context.Context, validator *validator.Validate, log logger.ILogger, e *echo.Echo) {
	group := e.Group("/api/v1/article_by_url")
	group.POST("/", getArticleByUrl(validator, log, ctx))
}

func getArticleByUrl(validator *validator.Validate, log logger.ILogger, 
ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *dtos.RequestArticleDto
		if err:= c.Bind(&request); err!=nil{
			fmt.Println("THIS IS FUCKING FAILING", err)

			return err
		}
		fmt.Println("request", request)
		result, err:= mediatr.Send[dtos.RequestArticleDto, *models.Article](ctx, *request)
		if err!=nil{ 
			fmt.Println("This motherfucking mediatr", err)
			return err }
		return c.JSON(http.StatusOK, result)

	}
}