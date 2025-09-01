-- SQL script to create the users table
-- This table will store user information for the application

CREATE TABLE users (
    id SERIAL PRIMARY KEY, -- Unique identifier for each user
    username VARCHAR(50) NOT NULL UNIQUE, -- Username must be unique
    password VARCHAR(255) NOT NULL, -- Password for the user
    email VARCHAR(100) NOT NULL UNIQUE, -- Email must be unique
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for when the user was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Timestamp for when the user was last updated
);