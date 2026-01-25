package models

// CreateUserRequest represents the request payload for creating a new user.
type CreateUserRequest struct {
	Username string // Username of the new user
	Email    string // Email address of the new user
	Country  string // Country of the new user
}
