package services

import (
	"errors"
	"sync"
)

// Balance struct to hold balance information
type Balance struct {
	UserID string  // User ID associated with the balance
	Amount float64 // Current balance amount
}

// BalanceService struct to manage balance operations
type BalanceService struct {
	balances map[string]*Balance // Map to hold user balances
	mu       sync.RWMutex        // Mutex for concurrent access
}

// NewBalanceService creates a new BalanceService
func NewBalanceService() *BalanceService {
	return &BalanceService{
		balances: make(map[string]*Balance), // Initialize the balances map
	}
}

// UpdateBalance updates the balance for a user
func (bs *BalanceService) UpdateBalance(userID string, amount float64) error {
	bs.mu.Lock() // Lock for writing
	defer bs.mu.Unlock() // Ensure unlock after function execution

	// Check if the user balance exists
	if balance, exists := bs.balances[userID]; exists {
		balance.Amount += amount // Update existing balance
	} else {
		// If not exists, create a new balance entry
		bs.balances[userID] = &Balance{
			UserID: userID,
			Amount: amount,
		}
	}
	return nil
}

// GetBalance retrieves the current balance for a user
func (bs *BalanceService) GetBalance(userID string) (float64, error) {
	bs.mu.RLock() // Lock for reading
	defer bs.mu.RUnlock() // Ensure unlock after function execution

	// Check if the user balance exists
	if balance, exists := bs.balances[userID]; exists {
		return balance.Amount, nil // Return the balance amount
	}
	return 0, errors.New("balance not found") // Return error if not found
}

// GetAllBalances retrieves all user balances
func (bs *BalanceService) GetAllBalances() map[string]float64 {
	bs.mu.RLock() // Lock for reading
	defer bs.mu.RUnlock() // Ensure unlock after function execution

	// Create a map to hold all balances
	allBalances := make(map[string]float64)
	for userID, balance := range bs.balances {
		allBalances[userID] = balance.Amount // Populate the map with user balances
	}
	return allBalances // Return all balances
}