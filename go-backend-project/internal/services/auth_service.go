package services

import (
	"errors"
	"fmt"
)

// User struct to represent a user in the system
type User struct {
	ID       int
	Username string
	Password string
	Role     string // e.g., "admin", "user"
}

// UserService struct to handle user-related operations
type UserService struct {
	users []User // In-memory user storage for simplicity
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{
		users: []User{},
	}
}

// Register a new user
func (us *UserService) RegisterUser(username, password, role string) (User, error) {
	// Check if the username already exists
	for _, user := range us.users {
		if user.Username == username {
			return User{}, errors.New("username already exists")
		}
	}

	// Create a new user
	newUser := User{
		ID:       len(us.users) + 1, // Simple ID assignment
		Username: username,
		Password: password, // In a real app, hash the password
		Role:     role,
	}

	// Add the new user to the in-memory storage
	us.users = append(us.users, newUser)
	return newUser, nil
}

// Login a user
func (us *UserService) LoginUser(username, password string) (User, error) {
	for _, user := range us.users {
		if user.Username == username && user.Password == password {
			return user, nil // Successful login
		}
	}
	return User{}, errors.New("invalid username or password")
}

// CheckUserRole checks if a user has a specific role
func (us *UserService) CheckUserRole(username, role string) (bool, error) {
	for _, user := range us.users {
		if user.Username == username {
			return user.Role == role, nil // Return true if roles match
		}
	}
	return false, fmt.Errorf("user %s not found", username)
}
