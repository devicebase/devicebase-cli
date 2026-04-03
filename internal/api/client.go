package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func NewClient() (*Client, error) {
	baseURL := strings.TrimRight(os.Getenv("DEVICEBASE_BASE_URL"), "/")
	if baseURL == "" {
		baseURL = "https://api.devicebase.cn"
	}

	apiKey := os.Getenv("DEVICEBASE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("DEVICEBASE_API_KEY environment variable is not set")
	}

	return &Client{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, c.BaseURL+path, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return data, fmt.Errorf("API error (HTTP %d): %s", resp.StatusCode, string(data))
	}

	return data, nil
}

func (c *Client) Post(path string, body interface{}) ([]byte, error) {
	req, err := c.newRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}
	return c.doRequest(req)
}

func (c *Client) Get(path string) ([]byte, error) {
	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	return c.doRequest(req)
}
