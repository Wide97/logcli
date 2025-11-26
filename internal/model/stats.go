package model

type Stats struct {
	Counts map[string]int `json:"counts"`
	Lines  int            `json:"lines"`
}
