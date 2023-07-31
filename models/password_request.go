package models

// PasswordRequest represents the request body structure.
type PasswordRequest struct {
	InitPassword string `json:"init_password"`
}
