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
	TTSStrategyHume       TTSStrategy = "hume"
	TTSStrategyMinimax    TTSStrategy = "minimax"
	TTSStrategyFish       TTSStrategy = "fish"
)

type TimeoutAction string

const (
	TimeoutActionEndCall      TimeoutAction = "end_call"
	TimeoutActionTransferCall TimeoutAction = "transfer_call"
)

type CallerTransformType string

const (
	CallerTransformAddPrefix       CallerTransformType = "add_prefix"
	CallerTransformAddSuffix       CallerTransformType = "add_suffix"
	CallerTransformReplace         CallerTransformType = "replace"
	CallerTransformStripDigitsEnd  CallerTransformType = "strip_digits_end"
	CallerTransformStripDigitsStart CallerTransformType = "strip_digits_start"
)

type SIPTrunkAuthType string

const (
	SIPTrunkAuthDigest SIPTrunkAuthType = "digest"
	SIPTrunkAuthACL    SIPTrunkAuthType = "acl"
	SIPTrunkAuthNone   SIPTrunkAuthType = "none"
)

type RegistrationStatus string

const (
	RegistrationNotRegistered RegistrationStatus = "not_registered"
	RegistrationPending       RegistrationStatus = "pending"
	RegistrationRegistered    RegistrationStatus = "registered"
	RegistrationFailed        RegistrationStatus = "failed"
	RegistrationAuthFailed    RegistrationStatus = "auth_failed"
)

type SipTrunkPBX string

const (
	SipTrunkPBX3CX     SipTrunkPBX = "3cx"
	SipTrunkPBXYeastar SipTrunkPBX = "yeastar"
	SipTrunkPBXWebex   SipTrunkPBX = "webex"
	SipTrunkPBXGeneric SipTrunkPBX = "generic"
)

type Organization struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type LLMStringsResponse struct {
	LLMs []string `json:"llms"`
}

type VoiceCatalogEntry struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	Provider    TTSStrategy  `json:"provider"`
	Languages   []string     `json:"languages"`
	PreviewURL  *string      `json:"preview_url"`
	Featured    *bool        `json:"featured,omitempty"`
}

type VoicesListResponse struct {
	Voices []VoiceCatalogEntry `json:"voices"`
}

type VoiceSettings struct {
	ID              string       `json:"id,omitempty"`
	Provider        TTSStrategy  `json:"provider,omitempty"`
	Speed           *float64     `json:"speed,omitempty"`
	Stability       *float64     `json:"stability,omitempty"`
	SimilarityBoost *float64     `json:"similarity_boost,omitempty"`
	Volume          *float64     `json:"volume,omitempty"`
}

type VadSettings struct {
	StartSecs          *float64 `json:"start_secs,omitempty"`
	StopSecs           *float64 `json:"stop_secs,omitempty"`
	SilenceTimeoutSecs *float64 `json:"silence_timeout_secs,omitempty"`
}

type NoiseCancelSettings struct {
	Enabled bool    `json:"enabled"`
	Level   float64 `json:"level"`
}

type TimeoutSettings struct {
	DefaultTimeout        *float64       `json:"default_timeout,omitempty"`
	TimeoutAction         *TimeoutAction `json:"timeout_action,omitempty"`
	TransferNumber        string         `json:"transfer_number,omitempty"`
	TransferMessage       string         `json:"transfer_message,omitempty"`
	SilenceCounter        *int           `json:"silence_counter,omitempty"`
	SilenceTimeoutAction  *TimeoutAction `json:"silence_timeout_action,omitempty"`
	SilenceTransferNumber string         `json:"silence_transfer_number,omitempty"`
}

type Agent struct {
	ID                           string               `json:"id"`
	PhoneNumberID                string               `json:"phone_number_id"`
	Name                         string               `json:"name"`
	LLM                          string               `json:"llm"`
	FallbackLLM                  string               `json:"fallback_llm"`
	Status                       AgentStatus          `json:"status"`
	Prompt                       string               `json:"prompt"`
	FirstMessage                 string               `json:"first_message"`
	FirstMessageInterruptionTime *float64            `json:"first_message_interruption_time,omitempty"`
	DefaultLanguage              string               `json:"default_language"`
	SummarizeCalls               bool                 `json:"summarize_calls"`
	CreatedAt                    string               `json:"created_at"`
	UpdatedAt                    string               `json:"updated_at"`
	NoiseCancelSettings          *NoiseCancelSettings `json:"noise_cancel_settings,omitempty"`
	VadSettings                  *VadSettings         `json:"vad_settings,omitempty"`
	EnableFirstMessageOutbound   bool                 `json:"enable_first_message_outbound"`
	VoiceSettings                *VoiceSettings       `json:"voice_settings,omitempty"`
	TimeoutSettings              *TimeoutSettings     `json:"timeout_settings,omitempty"`
	Keyterms                     []string             `json:"keyterms,omitempty"`
}

type AgentCreateInput struct {
	PhoneNumberID                *string              `json:"phone_number_id,omitempty"`
	Name                         string               `json:"name"`
	LLM                          string               `json:"llm,omitempty"`
	FallbackLLM                  string               `json:"fallback_llm,omitempty"`
	Status                       AgentStatus          `json:"status,omitempty"`
	Prompt                       string               `json:"prompt,omitempty"`
	FirstMessage                 string               `json:"first_message,omitempty"`
	FirstMessageInterruptionTime *float64            `json:"first_message_interruption_time,omitempty"`
	DefaultLanguage              string               `json:"default_language,omitempty"`
	SummarizeCalls               *bool                `json:"summarize_calls,omitempty"`
	NoiseCancelSettings          *NoiseCancelSettings `json:"noise_cancel_settings,omitempty"`
	VadSettings                  *VadSettings         `json:"vad_settings,omitempty"`
	EnableFirstMessageOutbound   *bool                `json:"enable_first_message_outbound,omitempty"`
	TimeoutSettings              *TimeoutSettings     `json:"timeout_settings,omitempty"`
	Keyterms                     []string             `json:"keyterms,omitempty"`
	VoiceSettings                *VoiceSettings       `json:"voice_settings,omitempty"`
}

type AgentUpdateInput struct {
	PhoneNumberID                *string              `json:"phone_number_id,omitempty"`
	Name                         string               `json:"name,omitempty"`
	LLM                          string               `json:"llm,omitempty"`
	FallbackLLM                  string               `json:"fallback_llm,omitempty"`
	Status                       AgentStatus          `json:"status,omitempty"`
	Prompt                       string               `json:"prompt,omitempty"`
	FirstMessage                 string               `json:"first_message,omitempty"`
	FirstMessageInterruptionTime *float64            `json:"first_message_interruption_time,omitempty"`
	DefaultLanguage              string               `json:"default_language,omitempty"`
	SummarizeCalls               *bool                `json:"summarize_calls,omitempty"`
	NoiseCancelSettings          *NoiseCancelSettings `json:"noise_cancel_settings,omitempty"`
	VadSettings                  *VadSettings         `json:"vad_settings,omitempty"`
	EnableFirstMessageOutbound   *bool                `json:"enable_first_message_outbound,omitempty"`
	TimeoutSettings              *TimeoutSettings     `json:"timeout_settings,omitempty"`
	Keyterms                     []string             `json:"keyterms,omitempty"`
	VoiceSettings                *VoiceSettings       `json:"voice_settings,omitempty"`
}

type GraphVersion struct {
	ID          string  `json:"id"`
	AgentID     string  `json:"agent_id"`
	CreatedAt   string  `json:"created_at"`
	PublishedAt *string `json:"published_at"`
	Tag         string  `json:"tag"`
}

type GraphNode map[string]interface{}

type GraphVersionDetail struct {
	Version GraphVersion `json:"version"`
	Nodes   []GraphNode  `json:"nodes"`
}

type DialAgentResponse struct {
	Success bool `json:"success"`
}

type TwilioPhoneNumber struct {
	ID           string `json:"id"`
	PhoneNumber  string `json:"phone_number"`
	PhoneSid     string `json:"phone_sid"`
	Label        string `json:"label"`
	AccountSid   string `json:"account_sid"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type TwilioPhoneNumberImportInput struct {
	Label        string `json:"label,omitempty"`
	PhoneNumber  string `json:"phone_number"`
	AccountSid   string `json:"account_sid"`
	AuthToken    string `json:"auth_token"`
}

type TwilioPhoneNumberLabelPatchInput struct {
	Label string `json:"label"`
}

type SipTrunkCallerIDTransform struct {
	Type   CallerTransformType `json:"type"`
	Value  string              `json:"value,omitempty"`
	Value2 string              `json:"value2,omitempty"`
	Number *int                `json:"number,omitempty"`
}

type SipTrunk struct {
	ID                         string                   `json:"id"`
	CreatedAt                  string                   `json:"created_at"`
	PBX                        SipTrunkPBX              `json:"pbx"`
	CallerIDTransforms         []SipTrunkCallerIDTransform `json:"caller_id_transforms,omitempty"`
	InboundHost                string                   `json:"inbound_host,omitempty"`
	InboundAuthType            SIPTrunkAuthType         `json:"inbound_auth_type,omitempty"`
	InboundSIPUsername         string                   `json:"inbound_sip_username,omitempty"`
	InboundRegistrationStatus  RegistrationStatus       `json:"inbound_registration_status,omitempty"`
	OutboundHost               string                   `json:"outbound_host,omitempty"`
	OutboundSIPUsername        string                   `json:"outbound_sip_username,omitempty"`
	Register                   bool                     `json:"register"`
	OutboundRegistrationStatus RegistrationStatus       `json:"outbound_registration_status,omitempty"`
	RemainInDialog             bool                     `json:"remain_in_dialog"`
}

type SipTrunkCreateInput struct {
	PBX                SipTrunkPBX               `json:"pbx"`
	CallerIDTransforms []SipTrunkCallerIDTransform `json:"caller_id_transforms,omitempty"`
	InboundHost        *string                   `json:"inbound_host,omitempty"`
	InboundAuthType    *SIPTrunkAuthType         `json:"inbound_auth_type,omitempty"`
	InboundSIPUsername *string                   `json:"inbound_sip_username,omitempty"`
	InboundSIPPassword *string                   `json:"inbound_sip_password,omitempty"`
	OutboundHost       *string                   `json:"outbound_host,omitempty"`
	OutboundSIPUsername *string                  `json:"outbound_sip_username,omitempty"`
	OutboundSIPPassword *string                  `json:"outbound_sip_password,omitempty"`
	Register           *bool                     `json:"register,omitempty"`
	RemainInDialog     *bool                     `json:"remain_in_dialog,omitempty"`
}

type SipTrunkUpdateInput struct {
	PBX                *SipTrunkPBX                `json:"pbx,omitempty"`
	CallerIDTransforms *[]SipTrunkCallerIDTransform `json:"caller_id_transforms,omitempty"`
	InboundHost        *string                     `json:"inbound_host,omitempty"`
	InboundAuthType    *SIPTrunkAuthType           `json:"inbound_auth_type,omitempty"`
	InboundSIPUsername *string                     `json:"inbound_sip_username,omitempty"`
	InboundSIPPassword *string                     `json:"inbound_sip_password,omitempty"`
	OutboundHost       *string                     `json:"outbound_host,omitempty"`
	OutboundSIPUsername *string                    `json:"outbound_sip_username,omitempty"`
	OutboundSIPPassword *string                    `json:"outbound_sip_password,omitempty"`
	Register           *bool                       `json:"register,omitempty"`
	RemainInDialog     *bool                       `json:"remain_in_dialog,omitempty"`
}

type SipTrunkAgentAssignment struct {
	ID              string `json:"id"`
	AgentID         string `json:"agent_id"`
	SipTrunkID      string `json:"sip_trunk_id"`
	Number          string `json:"number,omitempty"`
	AlternateNumber string `json:"alternate_number,omitempty"`
	CreatedAt       string `json:"created_at"`
}

type SipTrunkAgentAssignmentCreateInput struct {
	AgentID         string `json:"agent_id"`
	Number          string `json:"number,omitempty"`
	AlternateNumber string `json:"alternate_number,omitempty"`
}

type SessionTokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
