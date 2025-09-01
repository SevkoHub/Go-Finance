package main

import (
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "go-backend-project/pkg/logger"
    "go-backend-project/internal/database"
    "go-backend-project/internal/services"
    "github.com/go-chi/chi/v5"
)

// main is the entry point of the application.
func main() {
    // Initialize the logger
    log := logger.NewLogger()

    // Load the configuration (like environment variables)
    err := loadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Connect to the database
    db, err := database.Connect()
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    defer db.Close() // Ensure the database connection is closed when done

    // Set up the router
    r := chi.NewRouter()

    // Set up routes (for example, user and transaction routes)
    setupRoutes(r)

    // Start the HTTP server
    srv := &http.Server{
        Addr:    ":8080", // Port to listen on
        Handler: r,
    }

    // Channel to listen for interrupt signals
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    // Run the server in a goroutine
    go func() {
        log.Println("Starting server on :8080")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("ListenAndServe: %s", err)
        }
    }()

    // Wait for interrupt signal
    <-quit
    log.Println("Shutting down server...")

    // Create a context with a timeout for graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Shutdown the server gracefully
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }
    log.Println("Server exiting")
}

// loadConfig loads the configuration settings for the application.
func loadConfig() error {
    // Here you would load environment variables or config files
    // For simplicity, we are just returning nil
    return nil
}

// setupRoutes sets up the HTTP routes for the application.
func setupRoutes(r *chi.Mux) {
    // Example route for user registration
    r.Post("/auth/register", services.RegisterUser)
    // Example route for user login
    r.Post("/auth/login", services.LoginUser)
    // Add more routes as needed
}