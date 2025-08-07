package commands
import (
	"context"
	"encoding/json"

	"github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"


)
type GenerateDescriptionHandler struct {
	log               logger.ILogger
	rabbitmqPublisher rabbitmq.IPublisher
	// GenerateDescription contracts.ArticleRepository
	ctx               context.Context
	grpcClient        grpc.GrpcClient
}

func NewGenerateDescriptionHandler(log logger.ILogger, rabbitmqPublisher *rabbitmq.IPublisher, ctx context.Context, grpcClient grpc.GrpcClient) *GenerateDescriptionHandler {
	return &GenerateDescriptionHandler{log: log, rabbitmqPublisher: *rabbitmqPublisher, ctx: ctx, grpcClient: grpcClient}
}


func (dh *GenerateDescriptionHandler) Handle (){


	
}