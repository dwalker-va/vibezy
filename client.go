package vibezy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	baseURL = "https://app.officevibe.com/api/v2"

	//OfficeVibe's API returns an entire HTML login page if your apiKey is wrong, which will cause this error
	decodingErrorHint = "could not decode OfficeVibe response, have you tested whether your API key is set up correctly?: https://api.officevibe.com/docs/ping"
)

// NewClient returns a client for interacting with the OfficeVibe v2 API
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: baseURL,
		http:    &http.Client{},
	}
}

func apiError(status int, message string) error {
	return fmt.Errorf("OfficeVibe error: status: `%d`, message: `%s`", status, message)
}

// Client communicates with the OfficeVibe v2 API over HTTP using JSON
// You should use the `NewClient` constructor to create a new instance of this struct
type Client struct {
	apiKey  string
	baseURL string
	http    *http.Client
}

func (c *Client) buildRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.baseURL, path), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("content-type", "application/json")
	return req, nil
}

func (c *Client) doRequest(req *http.Request, format interface{}) (*http.Response, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(format)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", decodingErrorHint, err)
	}

	return resp, nil
}

// Ping calls the OfficeVibe v2 Ping API.
// This is useful to test whether your configuration (including apiKey) is working correctly.
func (c *Client) Ping(ctx context.Context) (*PingResponse, error) {
	req, err := c.buildRequest(ctx, http.MethodGet, "ping", nil)
	if err != nil {
		return nil, err
	}

	r := &PingResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || !r.IsSuccess {
		return r, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, nil
}

// ListUsers calls the OfficeVibe v2 users:list API.
func (c *Client) ListUsers(ctx context.Context) (*ListUsersResponse, error) {
	req, err := c.buildRequest(ctx, http.MethodGet, "users", nil)
	if err != nil {
		return nil, err
	}

	r := &ListUsersResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, err
}

// GetUser calls the OfficeVibe v2 users:get API.
func (c *Client) GetUser(ctx context.Context, email string) (*GetUserResponse, error) {
	req, err := c.buildRequest(ctx, http.MethodGet, fmt.Sprintf("users/%s", email), nil)
	if err != nil {
		return nil, err
	}

	r := &GetUserResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, err
}

// Update calls the OfficeVibe v2 users:update API.
// If a user does not already exist, they will be created and receive an invitation.
func (c *Client) UpdateUser(ctx context.Context, email string) (*UpdateUserResponse, error) {
	req, err := c.buildRequest(ctx, http.MethodPost, fmt.Sprintf("users/%s", email), nil)
	if err != nil {
		return nil, err
	}

	r := &UpdateUserResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, err
}

// Update calls the OfficeVibe v2 users:update API.
// If a user does not already exist, they will be created and receive an invitation.
func (c *Client) DeactivateUser(ctx context.Context, request DeactivateUserRequest) (*DeactivateUserResponse, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	req, err := c.buildRequest(ctx, http.MethodPost, "users/deactivate", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r := &DeactivateUserResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, err
}

// ListGroups calls the OfficeVibe v2 groups:list API.
func (c *Client) ListGroups(ctx context.Context) (*ListGroupsResponse, error) {
	req, err := c.buildRequest(ctx, http.MethodGet, "groups", nil)
	if err != nil {
		return nil, err
	}

	r := &ListGroupsResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, err
}

// GetGroup calls the OfficeVibe v2 groups:get API.
func (c *Client) GetGroup(ctx context.Context, groupID string) (*GetGroupResponse, error) {
	req, err := c.buildRequest(ctx, http.MethodGet, fmt.Sprintf("groups/?groupId=%s", groupID), nil)
	if err != nil {
		return nil, err
	}

	r := &GetGroupResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, err
}

// CreateGroup calls the OfficeVibe v2 group:create API.
func (c *Client) CreateGroup(ctx context.Context, request CreateGroupRequest) (*CreateGroupResponse, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	req, err := c.buildRequest(ctx, http.MethodPost, "groups", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r := &CreateGroupResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, err
}

// RemoveGroup calls the OfficeVibe v2 group:remove API.
func (c *Client) RemoveGroup(ctx context.Context, request RemoveGroupRequest) (*RemoveGroupResponse, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	req, err := c.buildRequest(ctx, http.MethodPost, "groups/remove", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r := &RemoveGroupResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, err
}

// AddUsersToGroup calls the OfficeVibe v2 group:addUsers API.
func (c *Client) AddUsersToGroup(ctx context.Context, request AddUsersToGroupRequest) (*AddUsersToGroupResponse, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	req, err := c.buildRequest(ctx, http.MethodPost, "groups/addUsers", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r := &AddUsersToGroupResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, err
}

// RemoveUsersFromGroup calls the OfficeVibe v2 group:removeUsers API.
func (c *Client) RemoveUsersFromGroup(ctx context.Context, request RemoveUsersFromGroupRequest) (*RemoveUsersFromGroupResponse, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	req, err := c.buildRequest(ctx, http.MethodPost, "groups/removeUsers", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r := &RemoveUsersFromGroupResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, err
}

// RemoveAllUsersFromGroup calls the OfficeVibe v2 group:removeAllUsers API.
func (c *Client) RemoveAllUsersFromGroup(ctx context.Context, request RemoveAllUsersFromGroupRequest) (*RemoveAllUsersFromGroupResponse, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	req, err := c.buildRequest(ctx, http.MethodPost, "groups/removeAllUsers", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r := &RemoveAllUsersFromGroupResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, r.ErrorMessage)
	}

	return r, err
}

// RemoveAllUsersFromGroup calls the OfficeVibe v2 group:removeAllUsers API.
func (c *Client) Sync(ctx context.Context, request SyncRequest) (*SyncResponse, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	req, err := c.buildRequest(ctx, http.MethodPost, "sync", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r := &SyncResponse{}
	resp, err := c.doRequest(req, r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || !r.IsSuccess {
		return nil, apiError(resp.StatusCode, strings.Join(r.Errors, ", "))
	}

	return r, err
}
