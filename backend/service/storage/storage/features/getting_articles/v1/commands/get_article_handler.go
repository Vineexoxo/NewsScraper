package commands

import (
	"context"

	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/pkg/utils"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/data/contracts"
	dtosv1 "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/dtos"
)

type GetArticlesHandler struct {
	log logger.ILogger
	rabbitmqPublisher *rabbitmq.IPublisher
	productRepository contracts.ArticleRepository 
	ctx context.Context
}


func NewGetArticlesHandler(log logger.ILogger, rabbitmqPublisher *rabbitmq.IPublisher, 
productRepository contracts.ArticleRepository, ctx context.Context) *GetArticlesHandler {
	return &GetArticlesHandler{log: log, rabbitmqPublisher: rabbitmqPublisher, productRepository: productRepository, ctx: ctx}
}


func(c* GetArticlesHandler) Handle(ctx context.Context, query *GetArticles)(*dtosv1.ResponseArticleDto, error){
	articles, err:= c.productRepository.GetAllArticles(ctx, query.ListQuery)
	if err!=nil{ return nil, err}
	_, err = utils.ListResultToListResultDto[*dtosv1.RequestArticleDto](articles)
	
	if err!=nil { return nil, err}
	return &dtosv1.ResponseArticleDto{}, nil
}