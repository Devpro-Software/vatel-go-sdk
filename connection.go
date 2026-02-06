package vatel

import (
	"context"
	"encoding/base64"
	"sync"

	"github.com/gorilla/websocket"
)

// Connection is a WebSocket connection to the call session. Send input audio and tool call outputs; receive server events via Receive or Messages().
type Connection struct {
	conn   *websocket.Conn
	mu     sync.Mutex
	closed bool
}

// DialConnection opens a WebSocket connection to wsURL (e.g. from Client.ConnectionURL(token)). Use ConnectionOptions to customize the dialer.
func DialConnection(ctx context.Context, wsURL string, opts *ConnectionOptions) (*Connection, error) {
	dialer := websocket.DefaultDialer
	if opts != nil && opts.Dialer != nil {
		dialer = opts.Dialer
	}
	conn, _, err := dialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		return nil, err
	}
	return &Connection{conn: conn}, nil
}

// ConnectionOptions holds optional settings for DialConnection.
type ConnectionOptions struct {
	Dialer *websocket.Dialer
}

// SendInputAudio sends a base64-encoded PCM 16 24kHz mono audio chunk to the server.
func (c *Connection) SendInputAudio(pcmBase64 string) error {
	return c.send(NewInputAudioMessage(pcmBase64))
}

// SendInputAudioBytes sends raw PCM bytes (encoded as base64) to the server.
func (c *Connection) SendInputAudioBytes(pcm []byte) error {
	return c.SendInputAudio(base64.StdEncoding.EncodeToString(pcm))
}

// SendToolCallOutput sends the result of a tool call requested by the server (see ServerMessage type "tool_call").
func (c *Connection) SendToolCallOutput(toolCallID, output string) error {
	return c.send(NewToolCallOutputMessage(toolCallID, output))
}

func (c *Connection) send(payload interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return ErrConnectionClosed
	}
	return c.conn.WriteJSON(payload)
}

// Receive reads the next server message. Use ServerMessage.ParseData() to get typed payloads by message type.
func (c *Connection) Receive() (ServerMessage, error) {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return ServerMessage{}, ErrConnectionClosed
	}
	c.mu.Unlock()

	var m ServerMessage
	err := c.conn.ReadJSON(&m)
	return m, err
}

// Messages returns a channel that receives all server messages until the connection is closed. The channel is closed when the read loop exits.
func (c *Connection) Messages() <-chan ServerMessage {
	ch := make(chan ServerMessage)
	go func() {
		defer close(ch)
		for {
			m, err := c.Receive()
			if err != nil {
				return
			}
			ch <- m
		}
	}()
	return ch
}

// Close closes the WebSocket connection with a normal closure.
func (c *Connection) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return nil
	}
	c.closed = true
	return c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

// CloseWithReason closes the WebSocket with the given code and reason text.
func (c *Connection) CloseWithReason(code int, text string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return nil
	}
	c.closed = true
	return c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, text))
}
