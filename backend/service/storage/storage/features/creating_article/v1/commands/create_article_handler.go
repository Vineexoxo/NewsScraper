package commands

import (
	"context"
	"encoding/json"

	"github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/data/contracts"

	dtosv1 "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/creating_article/v1/dtos"
)

type createArticleHandler struct {
	log logger.ILogger
	rabbitmqPublisher rabbitmq.IPublisher
	articleRepository contracts.ArticleRepository
	ctx context.Context
	grpcClient grpc.GrpcClient
}

func NewCreateArticleHandler(log logger.ILogger, rabbitmqPublisher *rabbitmq.IPublisher, articleRepository contracts.ArticleRepository, ctx context.Context, grpcClient grpc.GrpcClient) *createArticleHandler {
	return &createArticleHandler{log: log, rabbitmqPublisher: *rabbitmqPublisher, articleRepository: articleRepository, ctx: ctx, grpcClient: grpcClient}
}

func (c* createArticleHandler) Handle(ctx context.Context, command *CreateArticle) (*dtosv1.CreateArticleResponsetDto, error) {
	article := &models.Article{
		ArticleID:      command.ArticleID,
		Title:          command.Title,
		Link:           command.Link,
		Keywords:       command.Keywords,
		Creator:        command.Creator,
		Description:    command.Description,
		Content:        command.Content,
		PubDate:        command.PubDate,
		PubDateTZ:      command.PubDateTZ,
		ImageURL:       command.ImageURL,
		VideoURL:       command.VideoURL,
		SourceID:       command.SourceID,
		SourceName:     command.SourceName,
		SourcePriority: command.SourcePriority,
		SourceURL:      command.SourceURL,
		SourceIcon:     command.SourceIcon,
		Language:       command.Language,
		Country:        command.Country,
		Category:       command.Category,
		Sentiment:      command.Sentiment,
		SentimentStats: command.SentimentStats,
		AITag:          command.AITag,
		AIRegion:       command.AIRegion,
		AIOrg:          command.AIOrg,
		AISummary:      command.AISummary,
		AIContent:      command.AIContent,
		Duplicate:      command.Duplicate,
	}



	createdArticle, err:= c.articleRepository.CreateArticle(ctx, article)
	if err != nil {return nil, err}
	

	response := &dtosv1.CreateArticleResponsetDto{URL: createdArticle.Link}
	bytes, _ := json.Marshal(response)

	c.log.Info("CreateProductResponseDto", string(bytes))

	return response, nil
}

