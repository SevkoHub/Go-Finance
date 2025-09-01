-- SQL script to create the balances table
-- This table will store the balance information for users

CREATE TABLE balances (
    id SERIAL PRIMARY KEY, -- Unique identifier for each balance record
    user_id INT NOT NULL, -- Foreign key referencing the user
    amount DECIMAL(10, 2) NOT NULL, -- The balance amount, with two decimal places
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for when the balance was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp for when the balance was last updated
    FOREIGN KEY (user_id) REFERENCES users(id) -- Ensures that the user_id exists in the users table
);