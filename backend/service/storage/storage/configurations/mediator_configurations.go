package configurations

import (
	"context"
	"fmt"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/models"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/data/contracts"
	creating_article "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/creating_article/v1/commands"
	creating_article_dto "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/creating_article/v1/dtos"
	gettingArticlesUrl "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/get_article_by_url/commands"
	dtosGettingArticlesUrl "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/get_article_by_url/dtos"
	gettingArticles "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/commands"
	ResponseArticleDto "github.com/shishir54234/NewsScraper/backend/service/storage/storage/features/getting_articles/v1/dtos"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/grpc_client"
)

func ConfigArticlesMediator(log logger.ILogger, rabbitmqPublisher rabbitmq.IPublisher,
	articleRepository contracts.ArticleRepository,
	web_scraper_client grpcclient.WebScraperClient,
	ctx context.Context, grpcClient grpc.GrpcClient) error {

	// //https://stackoverflow.com/questions/72034479/how-to-implement-generic-interfaces
	// err := mediatr.RegisterRequestHandler[*creatingarticlev1commands.CreateArticle, *creatingarticlev1dtos.CreateArticleResponseDto](creatingarticlev1commands.NewCreateArticleHandler(log, rabbitmqPublisher, articleRepository, ctx, grpcClient))
	// if err != nil {
	// 	return err
	// }

	err := mediatr.RegisterRequestHandler[*gettingArticles.GetArticles,
		[]*ResponseArticleDto.ResponseArticleDto](gettingArticles.
		NewGetArticlesHandler(log, &rabbitmqPublisher, articleRepository, ctx))


	
	if err != nil {
		fmt.Println("Registering a request Handler for getting articles didnt work that well")
		return err
	}

	err = mediatr.RegisterRequestHandler[dtosGettingArticlesUrl.RequestArticleDto, *models.Article](
		gettingArticlesUrl.NewGetArticlesByUrlHandler(log, &rabbitmqPublisher, 
		articleRepository,web_scraper_client, ctx))
	

	if err != nil {
		fmt.Println("Registering a request Handler for getting articles by url didnt work that well")
		return err
	}


	err=mediatr.RegisterRequestHandler[*creating_article.CreateArticle, *creating_article_dto.CreateArticleResponsetDto](
		creating_article.NewCreateArticleHandler(log, &rabbitmqPublisher, articleRepository, ctx, grpcClient))
	if err!=nil {
		fmt.Println("Registering a request handler for creating articles didnt work as well as I wanted it to")
		return err
	}



	// err = mediatr.RegisterRequestHandler[*searchingarticlev1queries.SearchArticles, *searchingarticlev1dtos.SearchArticlesResponseDto](searchingarticlev1queries.NewSearchArticlesHandler(log, rabbitmqPublisher, articleRepository, ctx, grpcClient))
	// if err != nil {
	// 	return err
	// }

	// err = mediatr.RegisterRequestHandler[*updatingarticlev1commands.UpdateArticle, *updatingarticlev1dtos.UpdateArticleResponseDto](updatingarticlev1commands.NewUpdateArticleHandler(log, rabbitmqPublisher, articleRepository, ctx, grpcClient))
	// if err != nil {
	// 	return err
	// }

	// err = mediatr.RegisterRequestHandler[*deletingarticlev1commands.DeleteArticle, *mediatr.Unit](deletingarticlev1commands.NewDeleteArticleHandler(log, rabbitmqPublisher, articleRepository, ctx, grpcClient))
	// if err != nil {
	// 	return err
	// }

	// err = mediatr.RegisterRequestHandler[*gettingarticlebyidv1queries.GetArticleById, *gettingarticlebyidv1dtos.GetArticleByIdResponseDto](gettingarticlebyidv1queries.NewGetArticleByIdHandler(log, rabbitmqPublisher, articleRepository, ctx, grpcClient))
	
	if err != nil {
		return err
	}

	return nil
}
