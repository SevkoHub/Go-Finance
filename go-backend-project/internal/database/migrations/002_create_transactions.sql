-- SQL script to create the transactions table
-- This table will store all transaction records

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY, -- Unique identifier for each transaction
    user_id INT NOT NULL, -- ID of the user who made the transaction
    amount DECIMAL(10, 2) NOT NULL, -- Amount of money involved in the transaction
    transaction_type VARCHAR(10) NOT NULL, -- Type of transaction: 'credit' or 'debit'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the transaction was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp when the transaction was last updated
    FOREIGN KEY (user_id) REFERENCES users(id) -- Foreign key constraint linking to the users table
);