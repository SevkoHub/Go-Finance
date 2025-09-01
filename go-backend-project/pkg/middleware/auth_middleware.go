package middleware

import (
    "net/http"
    "strings"
)

// AuthMiddleware is a middleware function that checks if the request has a valid authorization token.
// It simulates a simple token-based authentication.
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the Authorization header from the request
        authHeader := r.Header.Get("Authorization")

        // Check if the Authorization header is present
        if authHeader == "" {
            // If not present, return a 401 Unauthorized response
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Split the header to get the token (assuming "Bearer <token>")
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            // If the format is incorrect, return a 401 Unauthorized response
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Here we would normally validate the token (omitted for simplicity)
        token := parts[1]
        if token != "valid-token" { // Simulating a valid token check
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // If the token is valid, call the next handler
        next.ServeHTTP(w, r)
    })
}