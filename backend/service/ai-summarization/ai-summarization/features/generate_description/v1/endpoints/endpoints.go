package endpoints

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/features/generate_description/v1/commands"
	descriptionpb "github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/grpc_server/proto"

	"google.golang.org/grpc"
)

// MapGrpc registers the gRPC server for this feature
func MapGrpc(grpcServer *grpc.Server, validate *validator.Validate, log logger.ILogger,rabbitmq *rabbitmq.IPublisher, ctx context.Context) {
	descriptionpb.RegisterDescriptionServiceServer(grpcServer, commands.NewGenerateDescriptionHandler(log,rabbitmq, ctx))
}