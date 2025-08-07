package grpc_client
import (
	"context"
	"github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/grpc_server/protos/description.proto" // import from article service
)

type ArticleClient interface {
	GetArticle(ctx context.Context, id string) (*articlepb.Article, error)
}

type articleClient struct {
	client articlepb.ArticleServiceClient
}

func NewArticleClient(client articlepb.ArticleServiceClient) ArticleClient {
	return &articleClient{client}
}

func (a *articleClient) GetArticle(ctx context.Context, id string) (*articlepb.Article, error) {
	req := &articlepb.GetArticleRequest{ArticleId: id}
	return a.client.GetArticle(ctx, req)
}