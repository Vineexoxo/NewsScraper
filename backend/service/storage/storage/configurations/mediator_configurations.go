package configurations

import (
	"context"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/data/contracts"
	gettingArticles "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/commands"	
	ResponseArticleDto "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/dtos"
)

func ConfigArticlesMediator(log logger.ILogger, rabbitmqPublisher rabbitmq.IPublisher,
	productRepository contracts.ArticleRepository, ctx context.Context, grpcClient grpc.GrpcClient) error {

	// //https://stackoverflow.com/questions/72034479/how-to-implement-generic-interfaces
	// err := mediatr.RegisterRequestHandler[*creatingproductv1commands.CreateArticle, *creatingproductv1dtos.CreateArticleResponseDto](creatingproductv1commands.NewCreateArticleHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	// if err != nil {
	// 	return err
	// }

	err := mediatr.RegisterRequestHandler[*gettingArticles.GetArticles,
		*ResponseArticleDto.ResponseArticleDto](gettingArticles.
		NewGetArticlesHandler(log, &rabbitmqPublisher, productRepository, ctx))
	// if err != nil {
	// 	return err
	// }

	// err = mediatr.RegisterRequestHandler[*searchingproductv1queries.SearchArticles, *searchingproductv1dtos.SearchArticlesResponseDto](searchingproductv1queries.NewSearchArticlesHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	// if err != nil {
	// 	return err
	// }

	// err = mediatr.RegisterRequestHandler[*updatingproductv1commands.UpdateArticle, *updatingproductv1dtos.UpdateArticleResponseDto](updatingproductv1commands.NewUpdateArticleHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	// if err != nil {
	// 	return err
	// }

	// err = mediatr.RegisterRequestHandler[*deletingproductv1commands.DeleteArticle, *mediatr.Unit](deletingproductv1commands.NewDeleteArticleHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	// if err != nil {
	// 	return err
	// }

	// err = mediatr.RegisterRequestHandler[*gettingproductbyidv1queries.GetArticleById, *gettingproductbyidv1dtos.GetArticleByIdResponseDto](gettingproductbyidv1queries.NewGetArticleByIdHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	
	if err != nil {
		return err
	}

	return nil
}
