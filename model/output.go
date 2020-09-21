package model

// Output represents the model for response
type Output struct {
	ID      int    `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
	Details error  `json:"details,omitempty"`
	Users   []User `json:"users,omitempty"`
}
