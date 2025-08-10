package commands

import (
	"context"
	"fmt"
	"reflect"

	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/pkg/utils"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/data/contracts"
	dtosv1 "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/dtos"
)

type GetArticlesHandler struct {
	log logger.ILogger
	rabbitmqPublisher *rabbitmq.IPublisher
	articleRepository contracts.ArticleRepository 
	ctx context.Context
}


func NewGetArticlesHandler(log logger.ILogger, rabbitmqPublisher *rabbitmq.IPublisher, 
articleRepository contracts.ArticleRepository, ctx context.Context) *GetArticlesHandler {
	return &GetArticlesHandler{log: log, rabbitmqPublisher: rabbitmqPublisher, articleRepository: articleRepository, ctx: ctx}
}


func(c* GetArticlesHandler) Handle(ctx context.Context, query *GetArticles)([]*dtosv1.ResponseArticleDto, error){
	
	articles, err:= c.articleRepository.GetAllArticles(ctx, query.ListQuery)
	if err!=nil || articles==nil { 
		fmt.Println("Error", err)
		fmt.Println("Articles", articles)
		return nil, err}
	fmt.Println("ARTICLE TYPE", reflect.TypeOf(articles.Items))
	
	ret, err := utils.ListResultToListResultDto[*dtosv1.ResponseArticleDto, *models.Article](articles)
	


	if err!=nil { return nil, err}
	return ret.Items, nil
}