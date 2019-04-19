package models

// Message structures..........
type Message struct {
	TraceID string                 `json:"trace_id"`
	Header  string                 `json:"header"`
	Payload map[string]interface{} `json:"payload"`
}
