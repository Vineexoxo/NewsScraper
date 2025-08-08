package main

import (
	"github.com/go-playground/validator"
	"github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/shishir54234/NewsScraper/backend/pkg/http"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/otel"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/configurations"
	"github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/llm_client"
	"github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/server"
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
				validator.New,
				llm_client.NewLLMClient,
			),
			fx.Invoke(server.RunServers),
		),
	).Run()

}