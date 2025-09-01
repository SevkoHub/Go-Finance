package middleware

import (
	"encoding/json"
	"net/http"
)

// ValidationMiddleware is a middleware that validates incoming requests
func ValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Here we can add validation logic for different endpoints
		// For example, we can check if the request method is POST for certain routes

		if r.Method != http.MethodPost {
			// If the method is not POST, we return a 405 Method Not Allowed
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// We can also validate the request body if needed
		var requestBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			// If there is an error decoding the JSON, we return a 400 Bad Request
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// If everything is fine, we call the next handler
		next.ServeHTTP(w, r)
	})
}