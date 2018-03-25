package model

// Auth model
type Auth struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
