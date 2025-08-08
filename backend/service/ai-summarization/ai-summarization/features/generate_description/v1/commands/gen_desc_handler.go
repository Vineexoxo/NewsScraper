package commands

import (
	"context"
	"fmt"

//	"github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
//	"github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/features/generate_description/v1/dtos"
	"github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/llm_client"
	descriptionpb "github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/grpc_server/proto"
//	dtosv1 "github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/features/generate_description/v1/dtos"
)
type GenerateDescriptionHandler struct {
	descriptionpb.UnimplementedDescriptionServiceServer
	log               logger.ILogger
	rabbitmqPublisher rabbitmq.IPublisher
	llmClient llm_client.LLMClient
	ctx               context.Context
}

func NewGenerateDescriptionHandler(log logger.ILogger, 
rabbitmqPublisher *rabbitmq.IPublisher, 
ctx context.Context) *GenerateDescriptionHandler {
	return &GenerateDescriptionHandler{log: log, 
	rabbitmqPublisher: *rabbitmqPublisher, 
	ctx: ctx}
}


func (dh *GenerateDescriptionHandler) GenerateDescription (ctx context.Context, 
query *descriptionpb.GenerateDescriptionRequest) (*descriptionpb.GenerateDescriptionResponse, error) {
	result, err:= dh.llmClient.GenerateDescription(ctx, query.Description)
	if err!=nil{
		fmt.Println("The llm client isnt working", err)
		return nil, err
	}
	return &descriptionpb.GenerateDescriptionResponse{
		Url: query.Url,
		Description: result,
	}, err
}

