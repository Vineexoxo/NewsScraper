package main

import (
	"github.com/go-playground/validator"
	"github.com/shishir54234/NewsScraper/backend/pkg/database"
	"github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/shishir54234/NewsScraper/backend/pkg/http"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/otel"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/configurations"
	"github.com/shishir54234/NewsScraper/backend/service/web-scraper/web-scraper/grpc_server/server"
	worker "github.com/shishir54234/NewsScraper/backend/service/web-scraper/worker/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				configurations.InitConfig,
				logger.InitLogger,
				http.NewContext,
				otel.TracerProvider,
				grpc.NewGrpcServer,
				rabbitmq.NewRabbitMQConn,
				rabbitmq.NewPublisher,
				func(cfg *configurations.Config) *database.RedisConfig {
					return cfg.Redis
				},
				func(cfg *database.RedisConfig) string {
					return database.NewRedisConnStr(cfg)
				},
				database.NewRedisDB,
				validator.New,
			),
			// Run the servers
			fx.Invoke(server.RunServers),
			fx.Invoke(worker.RunWorker),
		),
	).Run()
}
