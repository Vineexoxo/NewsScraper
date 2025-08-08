package server

import (
	"context"
	"errors"
	"net/http"

	grpc1 "github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/configurations"
	"github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/features/generate_description/v1/commands"
	"github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/llm_client"
	"go.uber.org/fx"

	// "google.golang.org/grpc"
	descriptionpb "github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/grpc_server/proto"
)



func RunServers(lc fx.Lifecycle,llm_client llm_client.LLMClient, log logger.ILogger, 
ctx context.Context, cfg *configurations.Config, rabbitmqPublisher rabbitmq.IPublisher, 
grpcServer *grpc1.GrpcServer) {
	lc.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				descriptionpb.RegisterDescriptionServiceServer(
				grpcServer.Grpc, // underlying *grpc.Server
				commands.NewGenerateDescriptionHandler(llm_client,log, &rabbitmqPublisher, ctx),
				)
				go func() {
					if err := grpcServer.RunGrpcServer(ctx); !errors.Is(err, http.ErrServerClosed) {
						log.Panicf("error running grpc server: %v", err)
					}
				}()
				return nil
			},
			OnStop: func(_ context.Context) error {
				log.Infof("all servers shutdown gracefully...")
				return nil
			},
	})
}