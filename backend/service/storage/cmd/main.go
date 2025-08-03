package main

import (
	"github.com/go-playground/validator"
	gormpgsql "github.com/shishir54234/NewsScraper/backend/pkg/database"
	"github.com/shishir54234/NewsScraper/backend/pkg/http"
	echoserver "github.com/shishir54234/NewsScraper/backend/pkg/http/echo/server"
	httpclient "github.com/shishir54234/NewsScraper/backend/pkg/httpclient"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/otel"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/storage/config"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/configurations"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/data/repositories"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/mappings"
	"github.com/shishir54234/NewsScraper/backend/service/storage/storage/models"
	"github.com/shishir54234/NewsScraper/backend/service/storage/server"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.InitConfig,
				logger.InitLogger,
				http.NewContext,
				echoserver.NewEchoServer,
				gormpgsql.NewPostgresDB,
				otel.TracerProvider,
				httpclient.New,
				repositories.NewPostgresArticleRepository,
				rabbitmq.NewRabbitMQConn,
				rabbitmq.NewPublisher,
				validator.New,
			),
			fx.Invoke(server.RunServers),
			fx.Invoke(func(gorm *gorm.DB) error {
				return gormpgsql.Migrate(gorm, &models.Product{})
			}),
			fx.Invoke(mappings.ConfigureMappings),
			fx.Invoke(configurations.ConfigEndpoints),
			fx.Invoke(configurations.ConfigArticlesMediator),
		),
	).Run()
}
