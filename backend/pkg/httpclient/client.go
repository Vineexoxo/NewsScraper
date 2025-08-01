package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	headers    map[string]string
}

type Option func(*Client)

// WithTimeout sets request timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithBaseURL sets base URL for all requests
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithHeaders sets default headers
func WithHeaders(headers map[string]string) Option {
	return func(c *Client) {
		c.headers = headers
	}
}

// New creates a new HTTP client
func New(opts ...Option) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		headers: make(map[string]string),
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Get performs GET request
func (c *Client) Get(ctx context.Context, path string) (*http.Response, error) {
	return c.do(ctx, "GET", path, nil)
}

// Post performs POST request
func (c *Client) Post(ctx context.Context, path string, body interface{}) (*http.Response, error) {
	return c.do(ctx, "POST", path, body)
}

// Put performs PUT request
func (c *Client) Put(ctx context.Context, path string, body interface{}) (*http.Response, error) {
	return c.do(ctx, "PUT", path, body)
}

// Delete performs DELETE request
func (c *Client) Delete(ctx context.Context, path string) (*http.Response, error) {
	return c.do(ctx, "DELETE", path, nil)
}

func (c *Client) do(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	url := c.baseURL + path

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

// DecodeJSON decodes JSON response
func DecodeJSON(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(v)
}