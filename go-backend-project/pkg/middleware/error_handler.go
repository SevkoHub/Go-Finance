package middleware

import (
	"net/http"
	"log"
)

// ErrorHandler is a middleware that handles errors in the application.
// It takes a http.Handler and returns a new http.Handler that wraps the original one.
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We use a defer function to recover from any panic that occurs in the handler
		defer func() {
			if err := recover(); err != nil {
				// Log the error message
				log.Printf("An error occurred: %v", err)
				// Respond with a 500 Internal Server Error
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}