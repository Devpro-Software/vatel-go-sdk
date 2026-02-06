package vatel

type AgentStatus string

const (
	AgentStatusActive   AgentStatus = "active"
	AgentStatusInactive AgentStatus = "inactive"
)

type TTSStrategy string

const (
	TTSStrategyElevenLabs TTSStrategy = "elevenlabs"
	TTSStrategyOpenAI     TTSStrategy = "openai"
	TTSStrategyCartesia   TTSStrategy = "cartesia"
)

type TimeoutAction string

const (
	TimeoutActionEndCall     TimeoutAction = "end_call"
	TimeoutActionTransferCall TimeoutAction = "transfer_call"
)

type Agent struct {
	ID                          string              `json:"id"`
	PhoneNumberID               string              `json:"phone_number_id"`
	Name                        string              `json:"name"`
	LLM                         string              `json:"llm"`
	Status                      AgentStatus         `json:"status"`
	Prompt                      string              `json:"prompt"`
	FirstMessage                string              `json:"first_message"`
	DefaultLanguage             string              `json:"default_language"`
	SummarizeCalls              bool                `json:"summarize_calls"`
	TTSStrategy                 TTSStrategy         `json:"tts_strategy"`
	CreatedAt                   string              `json:"created_at"`
	UpdatedAt                   string              `json:"updated_at"`
	NoiseCancelSettings         *NoiseCancelSettings `json:"noise_cancel_settings,omitempty"`
	VadSettings                 *VadSettings       `json:"vad_settings,omitempty"`
	EnableFirstMessageOutbound  bool                `json:"enable_first_message_outbound"`
	VoiceSettings               *VoiceSettings     `json:"voice_settings,omitempty"`
	TimeoutSettings             *TimeoutSettings   `json:"timeout_settings,omitempty"`
	Keyterms                    []string           `json:"keyterms,omitempty"`
}

type VoiceSettings struct {
	Speed            *float64 `json:"speed,omitempty"`
	Stability        *float64 `json:"stability,omitempty"`
	SimilarityBoost  *float64 `json:"similarity_boost,omitempty"`
	Volume           *float64 `json:"volume,omitempty"`
}

type VadSettings struct {
	StartSecs        *float64 `json:"start_secs,omitempty"`
	StartingSecs     *float64 `json:"starting_secs,omitempty"`
	StopSecs         *float64 `json:"stop_secs,omitempty"`
	StoppingSecs     *float64 `json:"stopping_secs,omitempty"`
	SilenceTimeoutSecs *float64 `json:"silence_timeout_secs,omitempty"`
}

type NoiseCancelSettings struct {
	Enabled bool    `json:"enabled"`
	Level   float64 `json:"level"`
}

type TimeoutSettings struct {
	DefaultTimeout        *float64      `json:"defaultTimeout,omitempty"`
	TimeoutAction         *TimeoutAction `json:"timeoutAction,omitempty"`
	TransferNumber        string        `json:"transferNumber,omitempty"`
	TransferMessage       string        `json:"transferMessage,omitempty"`
	SilenceCounter        *int          `json:"silenceCounter,omitempty"`
	SilenceTimeoutAction  *TimeoutAction `json:"silenceTimeoutAction,omitempty"`
	SilenceTransferNumber string        `json:"silenceTransferNumber,omitempty"`
	SilenceTransferMessage string       `json:"silenceTransferMessage,omitempty"`
}

type SessionTokenResponse struct {
	Token string `json:"token"`
}
