package models

type LLMResult struct {
	Intent   string   `json:"intent"`
	Entities []string `json:"entities"`
	Concepts []string `json:"concepts,omitempty"`
	Location string   `json:"location,omitempty"`
	RawQuery string   `json:"-"`
}
