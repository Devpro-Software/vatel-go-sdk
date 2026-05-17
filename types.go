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

type DialAgentOptions struct {
	Number       string
	Destination  string
	SipTrunkID   string
	CallerID     string
	FirstMessage string
	Prompt       string
}

type DialAgentResponse struct {
	Success bool `json:"success"`
}

type CallStatus string

const (
	CallStatusConnected  CallStatus = "connected"
	CallStatusStarted    CallStatus = "started"
	CallStatusInProgress CallStatus = "in_progress"
	CallStatusAuthFailed CallStatus = "auth_failed"
	CallStatusEnded      CallStatus = "ended"
)

type CallSource string

const (
	CallSourceTwilio     CallSource = "twilio"
	CallSourceSIP        CallSource = "sip"
	CallSourceSimulation CallSource = "simulation"
	CallSourceAPI        CallSource = "api"
)

type CallOutcome string

const (
	CallOutcomeTransferred  CallOutcome = "transferred"
	CallOutcomeEndedByAgent CallOutcome = "ended_by_agent"
	CallOutcomeEndedByUser  CallOutcome = "ended_by_user"
)

type ContextVariableType string

const (
	ContextVariableSystem    ContextVariableType = "system"
	ContextVariableInput     ContextVariableType = "input"
	ContextVariableExtracted ContextVariableType = "extracted"
)

type ContextVariableDataType string

const (
	ContextVariableDataTypeString  ContextVariableDataType = "string"
	ContextVariableDataTypeNumber  ContextVariableDataType = "number"
	ContextVariableDataTypeBoolean ContextVariableDataType = "boolean"
	ContextVariableDataTypeObject  ContextVariableDataType = "object"
	ContextVariableDataTypeArray   ContextVariableDataType = "array"
)

type ContextVariable struct {
	Name        string                  `json:"name,omitempty"`
	Type        ContextVariableType     `json:"type,omitempty"`
	Description string                  `json:"description,omitempty"`
	Value       interface{}             `json:"value,omitempty"`
	DataType    ContextVariableDataType `json:"dataType,omitempty"`
	Rationale   string                  `json:"rationale,omitempty"`
}

type TranscriptEntryType string

const (
	TranscriptEntryMessage          TranscriptEntryType = "message"
	TranscriptEntryToolCall         TranscriptEntryType = "tool_call"
	TranscriptEntryToolCallOutput   TranscriptEntryType = "tool_call_output"
	TranscriptEntryInterruption     TranscriptEntryType = "interruption"
)

type TranscriptToolCall struct {
	ItemID    string `json:"itemId,omitempty"`
	CallID    string `json:"callId,omitempty"`
	ToolName  string `json:"toolName,omitempty"`
	Arguments string `json:"arguments,omitempty"`
	Output    string `json:"output,omitempty"`
	StartedAt string `json:"startedAt,omitempty"`
	EndedAt   string `json:"endedAt,omitempty"`
}

type TranscriptEntry struct {
	Index            int                  `json:"index,omitempty"`
	Role             string               `json:"role,omitempty"`
	Message          string               `json:"message,omitempty"`
	Type             TranscriptEntryType  `json:"type,omitempty"`
	ToolCall         *TranscriptToolCall  `json:"toolCall,omitempty"`
	ToolCallOutput   string               `json:"toolCallOutput,omitempty"`
	CreatedAt        string               `json:"createdAt,omitempty"`
	DurationMs       int64                `json:"durationMs,omitempty"`
	TurnID           string               `json:"turnId,omitempty"`
}

type Transcript struct {
	Entries []TranscriptEntry `json:"entries"`
}

type CallFeedback struct {
	Stars int    `json:"stars"`
	Notes string `json:"notes"`
}

type Call struct {
	ID                  string             `json:"id"`
	AgentID             string             `json:"agent_id"`
	GraphVersionID      string             `json:"graph_version_id"`
	OrganizationID      string             `json:"organization_id"`
	PartyNumber         string             `json:"party_number"`
	Status              CallStatus         `json:"status"`
	Source              CallSource         `json:"source"`
	TerminationReason   string             `json:"termination_reason,omitempty"`
	ExtractedVariables  []ContextVariable  `json:"extracted_variables,omitempty"`
	Outbound            bool               `json:"outbound"`
	OutboundContactID   string             `json:"outbound_contact_id,omitempty"`
	OutboundListRunID   string             `json:"outbound_list_run_id,omitempty"`
	CreatedAt           string             `json:"created_at"`
	ConnectedAt         string             `json:"connected_at,omitempty"`
	StartedAt           string             `json:"started_at,omitempty"`
	EndedAt             string             `json:"ended_at,omitempty"`
	Summary             string             `json:"summary,omitempty"`
	Transcript          *Transcript        `json:"transcript,omitempty"`
	Cost                float64            `json:"cost,omitempty"`
	Tags                []string           `json:"tags,omitempty"`
	DurationSeconds     int                `json:"duration_seconds,omitempty"`
	Feedback            *CallFeedback      `json:"feedback,omitempty"`
}

type PaginationInfo struct {
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

type PaginatedCallsResponse struct {
	Calls      []Call         `json:"calls"`
	Pagination PaginationInfo `json:"pagination"`
}

type ListCallsParams struct {
	OrganizationID string
	Page           int
	PageSize       int
	AgentIDs       string
	Status         CallStatus
	Source         CallSource
	DateFrom       string
	DateTo         string
	Outbound       *bool
	Search         string
	Tag            string
	Outcome        CallOutcome
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

type SipTrunkAgentAssignmentPatchInput struct {
	Number          *string `json:"number,omitempty"`
	AlternateNumber *string `json:"alternate_number,omitempty"`
}

type SessionTransport string

const (
	SessionTransportWebSocket SessionTransport = "websocket"
	SessionTransportWebRTC    SessionTransport = "webrtc"
)

type GenerateSessionTokenRequest struct {
	AgentID      string            `json:"agent_id"`
	VersionID    string            `json:"version_id,omitempty"`
	ChatID       string            `json:"chat_id,omitempty"`
	Transport    *SessionTransport `json:"transport,omitempty"`
	FirstMessage string            `json:"first_message,omitempty"`
	Prompt       string            `json:"prompt,omitempty"`
	Chat         *bool             `json:"chat,omitempty"`
}

type SessionTokenResponse struct {
	Token       string  `json:"token"`
	Room        *string `json:"room,omitempty"`
	Identity    *string `json:"identity,omitempty"`
	URL         *string `json:"url,omitempty"`
	WebRTCToken *string `json:"webrtc_token,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
