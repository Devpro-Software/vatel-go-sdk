package vatel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

func (c *Client) doJSON(ctx context.Context, method, relPath string, query url.Values, reqBody interface{}, ok []int, decodeInto interface{}) error {
	base := strings.TrimSuffix(c.baseURL, "/")
	u, err := url.Parse(base + defaultRESTPath + relPath)
	if err != nil {
		return err
	}
	if query != nil {
		u.RawQuery = query.Encode()
	}

	var bodyReader io.Reader
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if decodeInto != nil {
		req.Header.Set("Accept", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if slices.Contains(ok, resp.StatusCode) {
		if decodeInto != nil && len(respBody) > 0 {
			if err := json.Unmarshal(respBody, decodeInto); err != nil {
				return fmt.Errorf("decode response: %w", err)
			}
		}
		return nil
	}
	return &APIError{StatusCode: resp.StatusCode, Body: respBody}
}

func (c *Client) doBytes(ctx context.Context, method, relPath string, query url.Values, accept string, ok []int) ([]byte, error) {
	base := strings.TrimSuffix(c.baseURL, "/")
	u, err := url.Parse(base + defaultRESTPath + relPath)
	if err != nil {
		return nil, err
	}
	if query != nil {
		u.RawQuery = query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	if accept != "" {
		req.Header.Set("Accept", accept)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if slices.Contains(ok, resp.StatusCode) {
		return respBody, nil
	}
	return nil, &APIError{StatusCode: resp.StatusCode, Body: respBody}
}
