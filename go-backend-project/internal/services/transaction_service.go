// internal/services/transaction_service.go

package services

import (
	"errors"
	"sync"
)

// Transaction struct represents a financial transaction
type Transaction struct {
	ID     string  // Unique identifier for the transaction
	Amount float64 // Amount of money involved in the transaction
	Type   string  // Type of transaction: "credit" or "debit"
	UserID string  // ID of the user associated with the transaction
}

// TransactionService struct provides methods for processing transactions
type TransactionService struct {
	mu sync.RWMutex // Mutex for handling concurrent access to transactions
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

// Credit method adds money to a user's account
func (ts *TransactionService) Credit(userID string, amount float64) error {
	ts.mu.Lock() // Locking for safe concurrent access
	defer ts.mu.Unlock() // Unlocking after the operation

	if amount <= 0 {
		return errors.New("amount must be greater than zero") // Validation check
	}

	// Here you would normally update the user's balance in the database
	// For simplicity, we just return nil to indicate success
	return nil
}

// Debit method removes money from a user's account
func (ts *TransactionService) Debit(userID string, amount float64) error {
	ts.mu.Lock() // Locking for safe concurrent access
	defer ts.mu.Unlock() // Unlocking after the operation

	if amount <= 0 {
		return errors.New("amount must be greater than zero") // Validation check
	}

	// Here you would normally check if the user has enough balance
	// For simplicity, we just return nil to indicate success
	return nil
}

// Transfer method moves money from one user to another
func (ts *TransactionService) Transfer(fromUserID string, toUserID string, amount float64) error {
	ts.mu.Lock() // Locking for safe concurrent access
	defer ts.mu.Unlock() // Unlocking after the operation

	if amount <= 0 {
		return errors.New("amount must be greater than zero") // Validation check
	}

	// Here you would normally check balances and perform the transfer
	// For simplicity, we just return nil to indicate success
	return nil
}

// History method retrieves transaction history for a user
func (ts *TransactionService) History(userID string) ([]Transaction, error) {
	// This would normally fetch transaction history from the database
	// For simplicity, we return an empty slice and nil error
	return []Transaction{}, nil
}