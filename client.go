// Package vatel provides a Go client for the Call Agent Builder REST and WebSocket APIs.
package vatel

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

const defaultRESTPath = "/v1"
const defaultWSPath = "/v1/connection"

// Client is the REST API client. Use New to construct it; call REST methods (e.g. ListAgents, GenerateSessionToken) for API access.
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

// GenerateSessionToken issues a short-lived credential for a voice session. Default transport is websocket when Transport is nil.
func (c *Client) GenerateSessionToken(ctx context.Context, in GenerateSessionTokenRequest) (*SessionTokenResponse, error) {
	var out SessionTokenResponse
	if err := c.doJSON(ctx, http.MethodPost, "/session-token", nil, in, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SessionToken returns a short-lived JWT for the given agent (websocket transport). Use the token with ConnectionURL or DialConnection.
func (c *Client) SessionToken(ctx context.Context, agentID string) (*SessionTokenResponse, error) {
	return c.GenerateSessionToken(ctx, GenerateSessionTokenRequest{AgentID: agentID})
}

// ListAgents returns all agents for the organization identified by the client's API key.
func (c *Client) ListAgents(ctx context.Context) ([]Agent, error) {
	var out []Agent
	if err := c.doJSON(ctx, http.MethodGet, "/agents", nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
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
