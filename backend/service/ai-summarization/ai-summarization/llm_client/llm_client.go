package llm_client

import (
	"bytes"
	"context"
	"encoding/json"
	// "fmt"
	"net/http"
)

type LLMClient interface {
	GenerateDescription(ctx context.Context, content string) (string, error)
}

type openaiClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

func NewLLMClient(apiKey, baseURL string) LLMClient {
	return &openaiClient{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		baseURL:    baseURL,
	}
}

func (o *openaiClient) GenerateDescription(ctx context.Context, content string) (string, error) {
	body, _ := json.Marshal(map[string]string{"content": content})

	req, _ := http.NewRequestWithContext(ctx, "POST", o.baseURL+"/v1/description", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+o.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := o.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Description string `json:"description"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Description, nil
}
