package model

// Login represents the model for an login
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
