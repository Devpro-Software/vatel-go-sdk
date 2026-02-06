package vatel

import (
	"errors"
	"fmt"
)

var ErrConnectionClosed = errors.New("vatel: connection closed")

type APIError struct {
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	if len(e.Body) > 0 {
		return fmt.Sprintf("vatel API error (status %d): %s", e.StatusCode, string(e.Body))
	}
	return fmt.Sprintf("vatel API error: status %d", e.StatusCode)
}
