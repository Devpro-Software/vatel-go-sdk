package vatel

import "encoding/json"

const (
	TypeSessionStarted       = "session_started"
	TypeResponseAudio        = "response_audio"
	TypeResponseText         = "response_text"
	TypeInputAudioTranscript = "input_audio_transcript"
	TypeSpeechStarted       = "speech_started"
	TypeSpeechStopped       = "speech_stopped"
	TypeSessionEnded        = "session_ended"
	TypeInterruption        = "interruption"
	TypeToolCall            = "tool_call"
	TypeInputAudio          = "input_audio"
	TypeToolCallOutput      = "tool_call_output"
)

// ServerMessage is a discriminated server event. Check Type and use ParseData() to get the typed payload.
type ServerMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data,omitempty"`
}

type SessionStartedData struct {
	ID string `json:"id"`
}

type ResponseAudioData struct {
	TurnID string `json:"turn_id"`
	Audio  string `json:"audio"`
}

type ResponseTextData struct {
	TurnID string `json:"turn_id"`
	Text   string `json:"text"`
}

type InputAudioTranscriptData struct {
	Transcript string `json:"transcript"`
}

type SpeechStartedData struct {
	Emulated bool `json:"emulated"`
}

type ToolCallArgument struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	DataType    string      `json:"dataType"`
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Value       interface{} `json:"value"`
}

type ToolCallData struct {
	ToolCallID string             `json:"toolCallId"`
	ToolName   string             `json:"toolName"`
	Arguments  []ToolCallArgument `json:"arguments"`
}

type InputAudioPayload struct {
	Type string `json:"type"`
	Data struct {
		Audio string `json:"audio"`
	} `json:"data"`
}

type ToolCallOutputPayload struct {
	Type string `json:"type"`
	Data struct {
		ToolCallID string `json:"toolCallId"`
		Output     string `json:"output"`
	} `json:"data"`
}

func NewInputAudioMessage(audioBase64 string) InputAudioPayload {
	return InputAudioPayload{
		Type: TypeInputAudio,
		Data: struct {
			Audio string `json:"audio"`
		}{Audio: audioBase64},
	}
}

func NewToolCallOutputMessage(toolCallID, output string) ToolCallOutputPayload {
	return ToolCallOutputPayload{
		Type: TypeToolCallOutput,
		Data: struct {
			ToolCallID string `json:"toolCallId"`
			Output     string `json:"output"`
		}{ToolCallID: toolCallID, Output: output},
	}
}

// ParseData unmarshals Data into the concrete type for this message (e.g. SessionStartedData, ResponseAudioData, ToolCallData). Returns nil for message types with no payload.
func (m *ServerMessage) ParseData() (interface{}, error) {
	switch m.Type {
	case TypeSessionStarted:
		var d SessionStartedData
		if err := json.Unmarshal(m.Data, &d); err != nil {
			return nil, err
		}
		return d, nil
	case TypeResponseAudio:
		var d ResponseAudioData
		if err := json.Unmarshal(m.Data, &d); err != nil {
			return nil, err
		}
		return d, nil
	case TypeResponseText:
		var d ResponseTextData
		if err := json.Unmarshal(m.Data, &d); err != nil {
			return nil, err
		}
		return d, nil
	case TypeInputAudioTranscript:
		var d InputAudioTranscriptData
		if err := json.Unmarshal(m.Data, &d); err != nil {
			return nil, err
		}
		return d, nil
	case TypeSpeechStarted:
		var d SpeechStartedData
		if err := json.Unmarshal(m.Data, &d); err != nil {
			return nil, err
		}
		return d, nil
	case TypeToolCall:
		var d ToolCallData
		if err := json.Unmarshal(m.Data, &d); err != nil {
			return nil, err
		}
		return d, nil
	case TypeSpeechStopped, TypeSessionEnded, TypeInterruption:
		return nil, nil
	default:
		return m.Data, nil
	}
}
