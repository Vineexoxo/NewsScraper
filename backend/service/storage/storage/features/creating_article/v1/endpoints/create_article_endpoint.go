package endpoints

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/creating_article/v1/commands"
	dtosv1 "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/creating_article/v1/dtos"
)

func MapRoute(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("/api/v1/articles/create")
	group.POST("/", createArticleHandler(validator, log, ctx))

}

func createArticleHandler(validator *validator.Validate, log logger.ILogger, ctx context.Context) echo.HandlerFunc {
	return func(c  echo.Context) error {
		var request *dtosv1.CreateArticleRequestDto
		if err:= c.Bind(&request); err!=nil{
			fmt.Println("THIS IS FUCKING FAILING", err)
			return err
		}
		command:= commands.NewCreateArticle(*request)

		result, err := mediatr.Send[*commands.CreateArticle, *dtosv1.CreateArticleResponsetDto](ctx, command)
		if err!=nil{ 
			fmt.Println("This motherfucking mediatr")
			return err }
		fmt.Println("result", result.URL)
		return nil
	}


}
