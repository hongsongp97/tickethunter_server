package model

type Comment struct {
	ID        uint   `json:"id,omitempty"`
	Owner     string `json:"owner,omitempty"`
	EventID   uint   `json:"eventID,omitempty"`
	Content   string `json:"content,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}
