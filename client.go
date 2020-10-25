package vibezy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	url = "https://app.officevibe.com/api/v2"

	//OfficeVibe's API returns an entire HTML login page if your apiKey is wrong, which will cause this error
	decodingErrorHint = "could not decode OfficeVibe response, have you tested whether your API key is set up correctly?: https://api.officevibe.com/docs/ping"
)

// NewClient returns a client for interacting with the OfficeVibe v2 API
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		http:   &http.Client{},
	}
}

// Client communicates with the OfficeVibe v2 API over HTTP using JSON
// You should use the `NewClient` constructor to create a new instance of this struct
type Client struct {
	apiKey string
	http   *http.Client
}

func (c *Client) buildRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", url, path), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("content-type", "application/json")
	return req, nil
}

// Ping calls the OfficeVibe v2 Ping API.
// This is useful to test whether your configuration (including apiKey) is working correctly.
func (c *Client) Ping(ctx context.Context) error {
	req, err := c.buildRequest(ctx, http.MethodGet, "ping", nil)
	if err != nil {
		return err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	type Response struct {
		IsSuccess    bool   `json:"isSuccess"`
		ErrorMessage string `json:"errorMessage"`
	}

	r := Response{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return fmt.Errorf("%s, %w", decodingErrorHint, err)
	}

	if resp.StatusCode != 200 || !r.IsSuccess {
		return fmt.Errorf("OfficeVibe error message: %s", r.ErrorMessage)
	}

	return nil
}
