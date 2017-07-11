package store

type command struct {
	Command string `json:"command,omitempty"`
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
}
