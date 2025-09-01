-- SQL script to create the audit_logs table
-- This table will store logs of actions performed in the application

CREATE TABLE audit_logs (
    id SERIAL PRIMARY KEY, -- Unique identifier for each log entry
    user_id INT NOT NULL, -- ID of the user who performed the action
    action VARCHAR(255) NOT NULL, -- Description of the action performed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of when the action was performed
    FOREIGN KEY (user_id) REFERENCES users(id) -- Foreign key reference to the users table
);