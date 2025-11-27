package model

type ErrorDetail struct {
	Line int    `json:"line"`
	Text string `json:"text"`
}
type Stats struct {
	Counts map[string]int `json:"counts"`
	Lines  int            `json:"lines"`
	Errors []ErrorDetail  `json:"errors,omitempty"`
}
