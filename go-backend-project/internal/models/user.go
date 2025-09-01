// internal/models/user.go

package models

import (
	"encoding/json"
	"errors"
)

// User struct represents a user in the system
type User struct {
	ID       int    `json:"id"`       // Unique identifier for the user
	Username string `json:"username"` // Username of the user
	Password string `json:"password"` // Password of the user (should be hashed in a real application)
	Role     string `json:"role"`     // Role of the user (e.g., admin, user)
}

// NewUser creates a new user with the given details
func NewUser(id int, username, password, role string) *User {
	return &User{
		ID:       id,
		Username: username,
		Password: password,
		Role:     role,
	}
}

// Validate checks if the user fields are valid
func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

// ToJSON converts the User struct to JSON format
func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}

// FromJSON populates the User struct from JSON data
func (u *User) FromJSON(data []byte) error {
	return json.Unmarshal(data, u)
}