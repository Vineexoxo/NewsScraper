package main

import (
	"github.com/go-playground/validator"
	gormpgsql "github.com/shishir54234/NewsScraper/backend/pkg/database"
	// "github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/shishir54234/NewsScraper/backend/pkg/http"
	echoserver "github.com/shishir54234/NewsScraper/backend/pkg/http/echo/server"
	httpclient "github.com/shishir54234/NewsScraper/backend/pkg/httpclient"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/otel"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/storage/config"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/configurations"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/data/repositories"
	grpcclient "github.com/shishir54234/NewsScraper/backend/service/storage/storage/grpc_client"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/mappings"
	
	// "github.com/shishir54234/NewsScraper/backend/pkg/models"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/server"
	"go.uber.org/fx"
	// "gorm.io/gorm"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.InitConfig,
				gormpgsql.NewPostGresConnStr,
				logger.InitLogger,
				http.NewContext,
				echoserver.NewEchoServer,
				gormpgsql.NewPostgresDB,
				otel.TracerProvider,
				httpclient.New,
				func () *grpcclient.WebScraperClient { return grpcclient.NewWebScraperClient("localhost:6600") },
				grpcclient.NewLLMClient,
				repositories.NewPostgresArticleRepository,
				rabbitmq.NewRabbitMQConn,
			
				
				validator.New,
				rabbitmq.NewPublisher,
			),
			fx.Invoke(server.RunServers),
			fx.Invoke(mappings.ConfigureMappings),
			fx.Invoke(configurations.ConfigEndpoints),
			fx.Invoke(configurations.ConfigArticlesMediator),
		),
	).Run()
}
