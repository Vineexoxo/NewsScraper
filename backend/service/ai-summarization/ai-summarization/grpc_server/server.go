
package grpc_server
import 	articlepb "github.com/shishir54234/NewsScraper/backend/service/ai-summarization/ai-summarization/grpc_server/proto"



type summarizationServer struct {
	articleClient articlepb.ArticleServiceClient  // <-- this is your gRPC client to article service
	// maybe also: llmClient, logger, etc.
}