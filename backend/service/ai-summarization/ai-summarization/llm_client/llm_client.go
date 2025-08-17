package llm_client

import (
	// "bytes"
	"context"
	"strings"
	// "encoding/json"
	"fmt"

	// "fmt"
	"net/http"

	"github.com/shishir54234/NewsScraper/backend/pkg/config"
	"google.golang.org/genai"
)

type LLMClient interface {
	GenerateDescription(ctx context.Context, content string) (string, []string, error)
}

type openaiClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

type geminiAiClient struct {
	client     *genai.Client
	apiKey     string
	baseURL    string
}


func NewLLMClient(llm_client_config *config.LlmConfig) LLMClient {

	config := genai.ClientConfig{
		APIKey:  llm_client_config.ApiKey,
		Backend: genai.BackendGeminiAPI,
	}
	new_client, err:= genai.NewClient(context.Background(), &config)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return &geminiAiClient{
		client:     new_client,
		apiKey:     llm_client_config.ApiKey,
		baseURL:    llm_client_config.BaseURL,
	}


}

func (g *geminiAiClient) GenerateDescription(ctx context.Context, content string) (string, []string, error) {
	result, err:= g.client.Models.GenerateContent(ctx, "gemini-2.0-flash",
        genai.Text("Summarise this in a easy to understand way but also in sufficient number of words and give it to me in a way that you define it in different lines be seperated by /n "+
		content),
        nil,
    )
	if err!=nil{
		fmt.Println("There is problem in getting summary", err)
	}
	keywords, err:= g.client.Models.GenerateContent(ctx, "gemini-2.0-flash",
		genai.Text("Give me an array of keywords related to this article in a easy to understand way but also in sufficient number of words let the words seperated by a comma, give at maximum 5 keywords  "+
		result.Text()),
		nil,
	)	
	if err!=nil{
		fmt.Println("There is problem in getting keywords", err)
	}
	res:= strings.Split(keywords.Text(), ",")	
    if err != nil {
        fmt.Println("Error in splitting keywords", err)
		return "", nil, err
    }
    fmt.Println(result.Text())
	return result.Text(),res, nil
}