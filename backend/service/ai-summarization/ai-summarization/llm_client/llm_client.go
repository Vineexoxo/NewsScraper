package llm_client

import (
	// "bytes"
	"context"
	"regexp"
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

func (g *geminiAiClient) cleanKeyword(keyword string) string {
	// Remove extra whitespace and common punctuation
	keyword = strings.TrimSpace(keyword)
	keyword = strings.Trim(keyword, ".,!?;:\"'")
	
	// Remove numbering (e.g., "1. Technology" -> "Technology")
	re := regexp.MustCompile(`^\d+\.\s*`)
	keyword = re.ReplaceAllString(keyword, "")
	
	// Remove bullet points and dashes
	keyword = strings.TrimLeft(keyword, "•-* ")
	keyword = strings.TrimSpace(keyword)
	
	return keyword
}
func (g* geminiAiClient) parseKeywords(keywordsText string) []string {
	// split by comma and clean each word
	rawKeywords:=strings.Split(keywordsText, ",")
	var cleanedKeywords []string

	for _,keyword:=range rawKeywords{
		cleanedKeyword:=g.cleanKeyword(keyword)
		if cleanedKeyword!="" && len(cleanedKeyword)>2{
			cleanedKeywords=append(cleanedKeywords, cleanedKeyword)
		}		
	}
	if len(cleanedKeywords) > 5 {
		cleanedKeywords = cleanedKeywords[:5]
	}
	return cleanedKeywords
} 
func (g* geminiAiClient) cleanText(text string) string {
	re := regexp.MustCompile(`\n{3,}`)
	text = re.ReplaceAllString(text, "\n\n")
	lines := strings.Split(text, "\n")
	var cleanedLines []string
	for _, line := range lines {
		cleanedLine := strings.TrimSpace(line)
		if cleanedLine != "" {
			cleanedLines = append(cleanedLines, cleanedLine)
		}
	}
	return strings.Join(cleanedLines, "\n")
}
func (g *geminiAiClient) generateSummary(ctx context.Context, content string)(string, error){
	prompt := fmt.Sprintf(`Please provide a clear, comprehensive summary of the following content. 
		The summary should:
		- Be 2-4 paragraphs long
		- Use clear, accessible language
		- Capture the main points and key information
		- Be well-structured with proper paragraph breaks
		- Focus on the most important and newsworthy aspects

		Content to summarize:
		%s`, content)

	result, err := g.client.Models.GenerateContent(ctx, "gemini-2.0-flash-exp",
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			Temperature:     genai.Ptr(float32(0.3)), // Lower temperature for more focused summaries
			MaxOutputTokens: *genai.Ptr(int32(500)),   // Limit output length
			TopP:            genai.Ptr(float32(0.8)),
		},
	)
	if err != nil {
		return "", err
	}
	summary:= strings.TrimSpace(result.Text())
	summary = g.cleanText(summary)
	return summary, nil
}

func (g *geminiAiClient) generateKeywords(ctx context.Context, content string, summary string) ([]string, error) {
	prompt:= fmt.Sprintf(`Extract 5 relevant keywords from this content. Return ONLY the keywords separated by commas, no other text.
			Rules:
			- Focus on main topics, entities, and themes
			- Use single words or short phrases (2-3 words max)
			- Avoid generic words like "article", "news", "report"
			- Prefer specific, searchable terms
			- Return exactly 5 keywords

			Content: %s

			Summary: %s`, content, summary)
	result,err:=g.client.Models.GenerateContent(ctx, "gemini-2.0-flash", genai.Text(prompt),&genai.GenerateContentConfig{
		Temperature:     genai.Ptr(float32(0.1)), // Lower temperature for more focused summaries
		MaxOutputTokens: *genai.Ptr(int32(100)),   // Limit output length
		TopP:            genai.Ptr(float32(0.8)),
	})
	if err != nil {
		return nil, err
	}
	if result == nil || result.Text() == "" {
		return nil, nil
	}
	keywordsText:=strings.TrimSpace(result.Text())
	keywords:=g.parseKeywords(keywordsText)
	return keywords, nil
}


func (g *geminiAiClient) GenerateDescription(ctx context.Context, content string) (string, []string, error) {
	result, err:= g.generateSummary(ctx, content)
	if err!=nil{
		fmt.Println("The llm client isnt working", err)
		return "",nil, err
	}
	keywords, err:= g.generateKeywords(ctx, content, result)
	if err!=nil{
		fmt.Println("The llm client isnt working", err)
		return "",nil, err
	}
	return result, keywords, nil
}