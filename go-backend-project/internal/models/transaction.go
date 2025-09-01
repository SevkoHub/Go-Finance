// internal/models/transaction.go

package models

import (
	"encoding/json"
	"errors"
	"time"
)

// Transaction struct represents a financial transaction.
// It contains fields for ID, Amount, Type (credit/debit), UserID, and Timestamp.
type Transaction struct {
	ID        int       `json:"id"`        // Unique identifier for the transaction
	Amount    float64   `json:"amount"`    // Amount of money involved in the transaction
	Type      string    `json:"type"`      // Type of transaction: "credit" or "debit"
	UserID    int       `json:"user_id"`   // ID of the user associated with the transaction
	Timestamp time.Time  `json:"timestamp"` // Time when the transaction occurred
}

// NewTransaction creates a new transaction instance.
// It validates the transaction type and sets the current timestamp.
func NewTransaction(amount float64, transactionType string, userID int) (*Transaction, error) {
	if transactionType != "credit" && transactionType != "debit" {
		return nil, errors.New("invalid transaction type") // Return error if type is not valid
	}

	return &Transaction{
		Amount:    amount,
		Type:      transactionType,
		UserID:    userID,
		Timestamp: time.Now(), // Set the current time as the transaction timestamp
	}, nil
}

// ToJSON converts the Transaction struct to JSON format.
// This is useful for sending transaction data over HTTP.
func (t *Transaction) ToJSON() ([]byte, error) {
	return json.Marshal(t) // Convert struct to JSON
}

// FromJSON populates the Transaction struct from JSON data.
// This is useful for reading transaction data from HTTP requests.
func (t *Transaction) FromJSON(data []byte) error {
	return json.Unmarshal(data, t) // Populate struct from JSON
}