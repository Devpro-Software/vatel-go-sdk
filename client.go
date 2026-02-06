// Package vatel provides a Go client for the Call Agent Builder REST and WebSocket APIs.
package vatel

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const defaultRESTPath = "/v1"
const defaultWSPath = "/v1/connection"

// Client is the REST API client. Use New to construct it; use SessionToken and ListAgents for API calls.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// ClientOption configures a Client (e.g. WithHTTPClient).
type ClientOption func(*Client)

// WithHTTPClient sets the HTTP client used for REST requests.
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = c
	}
}

// New builds a Client. baseURL can be a host (e.g. "api.vatel.ai") or full URL; apiKey is the organization Bearer token.
func New(baseURL, apiKey string, opts ...ClientOption) *Client {
	baseURL = strings.TrimSuffix(baseURL, "/")
	if baseURL != "" && !strings.HasPrefix(baseURL, "http") {
		baseURL = "https://" + baseURL
	}
	c := &Client{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

// SessionToken returns a short-lived JWT for the given agent. Use the token with ConnectionURL or DialConnection.
func (c *Client) SessionToken(ctx context.Context, agentID string) (*SessionTokenResponse, error) {
	u, err := url.Parse(c.baseURL + defaultRESTPath + "/session-token")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("agentId", agentID)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, &APIError{StatusCode: resp.StatusCode, Body: body}
	}

	var out SessionTokenResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("decode session token: %w", err)
	}
	return &out, nil
}

// ListAgents returns all agents for the organization identified by the client's API key.
func (c *Client) ListAgents(ctx context.Context) ([]Agent, error) {
	u := c.baseURL + defaultRESTPath + "/agents"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, &APIError{StatusCode: resp.StatusCode, Body: body}
	}

	var out []Agent
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("decode agents: %w", err)
	}
	return out, nil
}

// ConnectionURL returns the WebSocket URL for the connection channel with the given JWT (e.g. from SessionToken).
func (c *Client) ConnectionURL(token string) string {
	u := c.baseURL
	u = strings.Replace(u, "https://", "wss://", 1)
	u = strings.Replace(u, "http://", "ws://", 1)
	return u + defaultWSPath + "?token=" + url.QueryEscape(token)
}

// DialConnection opens a WebSocket connection using a session token obtained from SessionToken.
func (c *Client) DialConnection(ctx context.Context, token string) (*Connection, error) {
	return DialConnection(ctx, c.ConnectionURL(token), nil)
}
