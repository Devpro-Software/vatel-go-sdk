# Vatel Go SDK

Use the Vatel Call Agent Builder from Go: manage agents over the REST API and run real-time voice sessions over WebSockets.

## Installation

```bash
go get github.com/Devpro-Software/vatel-go-sdk
```

## Quick start

Create a client with your API base URL and organization API key, then open a call session:

```go
package main

import (
	"context"
	"log"

	"github.com/Devpro-Software/vatel-go-sdk"
)

func main() {
	client := vatel.New("https://api.vatel.ai", "your-org-api-key")
	ctx := context.Background()

	tokenResp, err := client.SessionToken(ctx, "agent-uuid")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := client.DialConnection(ctx, tokenResp.Token)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for msg := range conn.Messages() {
		data, _ := msg.ParseData()
		switch msg.Type {
		case vatel.TypeSessionStarted:
			if d, ok := data.(vatel.SessionStartedData); ok {
				log.Println("session started", d.ID)
			}
		case vatel.TypeResponseAudio:
			if d, ok := data.(vatel.ResponseAudioData); ok {
				// Decode d.Audio (base64) and play PCM 16 24kHz mono
				_ = d
			}
		case vatel.TypeToolCall:
			if d, ok := data.(vatel.ToolCallData); ok {
				conn.SendToolCallOutput(d.ToolCallID, "result")
			}
		case vatel.TypeSessionEnded:
			return
		}
	}
}
```

## REST API

Use the same client to call the REST API:

- **List agents** — `client.ListAgents(ctx)` returns all agents for your organization.
- **Session token** — `client.SessionToken(ctx, agentID)` returns a short-lived JWT. Use this token to open a WebSocket connection for that agent.

You can pass a custom HTTP client:

```go
client := vatel.New("https://api.vatel.ai", apiKey, vatel.WithHTTPClient(myHTTPClient))
```

## WebSocket (call session)

After you have a session token:

1. **Connect** — `client.DialConnection(ctx, tokenResp.Token)` opens the WebSocket. You can also build the URL with `client.ConnectionURL(token)` and call `vatel.DialConnection(ctx, url, nil)` if you need a custom dialer.
2. **Send** — Use `conn.SendInputAudio(base64PCM)` or `conn.SendInputAudioBytes(pcm)` for microphone audio (PCM 16-bit 24 kHz mono). Use `conn.SendToolCallOutput(toolCallID, output)` when the server sends a tool call.
3. **Receive** — Read one message with `conn.Receive()`, or consume all messages with `conn.Messages()`. Each message has a `Type` and optional `Data`; call `msg.ParseData()` to get typed structs such as `SessionStartedData`, `ResponseAudioData`, `ToolCallData`, etc.

When you’re done, call `conn.Close()` (or `conn.CloseWithReason(code, text)`).

## Error handling

- REST errors are returned as `*vatel.APIError` with `StatusCode` and `Body`. Check the status code and body to handle 4xx/5xx responses.
- If you send on a closed WebSocket connection, you get `vatel.ErrConnectionClosed`.
