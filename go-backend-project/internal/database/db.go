package database

import (
    "database/sql" // Go's database/sql package for database operations
    "log"          // Log package for logging errors
    _ "github.com/lib/pq" // PostgreSQL driver
    "os"           // Package for accessing environment variables
)

// DB is a global variable to hold the database connection
var DB *sql.DB

// Connect initializes the database connection
func Connect() {
    // Get the database URL from environment variables
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL environment variable is not set")
    }

    // Open a connection to the database
    var err error
    DB, err = sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal("Error opening database: ", err)
    }

    // Check if the database is reachable
    err = DB.Ping()
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }

    log.Println("Database connection established")
}

// Close closes the database connection
func Close() {
    if err := DB.Close(); err != nil {
        log.Fatal("Error closing the database: ", err)
    }
    log.Println("Database connection closed")
}

// Migrate runs the database migrations
func Migrate() {
    // Here we would typically run the SQL migration files
    // For simplicity, we will just log that migrations would run
    log.Println("Running database migrations...")
    // In a real application, you would read the SQL files and execute them
}