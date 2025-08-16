package grpcclient

import (
	"context"
	"fmt"
	"github.com/shishir54234/NewsScraper/backend/service/storage/config"
	descriptionpb "github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/grpc_server/proto"
	"google.golang.org/grpc"
)


type LLMClient struct {
	client descriptionpb.DescriptionServiceClient

}
func NewLLMClient(config *config.ConfigLLMClient) *LLMClient {
	fmt.Println("SHJOLJJJJ", config.ConnAddr)
	conn , err:= grpc.Dial(config.ConnAddr, grpc.WithInsecure())
	if err!=nil{
		fmt.Println("Failed to dial server:", err)
		return nil
	}

	client := descriptionpb.NewDescriptionServiceClient(conn)
	return &LLMClient{
		client: client,
	}
}
func (clt *LLMClient) GenerateDescription (ctx context.Context, content string) (string, error) {
	result, err:= clt.client.GenerateDescription(ctx, &descriptionpb.GenerateDescriptionRequest{Description: content})
	if err!=nil{
		fmt.Println("The llm client isnt working", err)
		return "", err
	}
	return result.Description, err
}