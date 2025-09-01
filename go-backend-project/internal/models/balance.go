// internal/models/balance.go

package models

import (
	"encoding/json" // JSON encoding and decoding
	"sync"          // For RWMutex
)

// Balance struct represents a user's balance
type Balance struct {
	UserID string  `json:"user_id"` // User ID associated with the balance
	Amount float64 `json:"amount"`   // Current balance amount
	mu     sync.RWMutex // Mutex for concurrent access
}

// NewBalance creates a new Balance instance
func NewBalance(userID string, initialAmount float64) *Balance {
	return &Balance{
		UserID: userID,
		Amount: initialAmount,
	}
}

// UpdateBalance updates the balance amount safely
func (b *Balance) UpdateBalance(amount float64) {
	b.mu.Lock() // Lock for writing
	defer b.mu.Unlock() // Ensure unlock after function completes
	b.Amount += amount // Update the balance
}

// GetBalance returns the current balance amount safely
func (b *Balance) GetBalance() float64 {
	b.mu.RLock() // Lock for reading
	defer b.mu.RUnlock() // Ensure unlock after function completes
	return b.Amount // Return the current balance
}

// MarshalJSON customizes the JSON encoding for Balance
func (b *Balance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserID string  `json:"user_id"`
		Amount float64 `json:"amount"`
	}{
		UserID: b.UserID,
		Amount: b.Amount,
	})
}

// UnmarshalJSON customizes the JSON decoding for Balance
func (b *Balance) UnmarshalJSON(data []byte) error {
	aux := &struct {
		UserID string  `json:"user_id"`
		Amount float64 `json:"amount"`
	}{}
	if err := json.Unmarshal(data, aux); err != nil {
		return err // Return error if unmarshalling fails
	}
	b.UserID = aux.UserID
	b.Amount = aux.Amount
	return nil // Return nil if successful
}