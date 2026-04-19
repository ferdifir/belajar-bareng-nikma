package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"nikma/internal/services"
)

// AuthMiddleware provides authentication middleware
type AuthMiddleware struct {
	authService *services.AuthService
}

// NewAuthMiddleware creates a new auth middleware
func NewAuthMiddleware(authService *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{authService: authService}
}

// CheckAuth checks if the request contains valid authentication
func (m *AuthMiddleware) CheckAuth(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	// Check if it's a basic auth header
	if strings.HasPrefix(authHeader, "Basic ") {
		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			return false
		}

		credentials := string(decodedCredentials)
		parts := strings.Split(credentials, ":")
		if len(parts) != 2 {
			return false
		}

		username := parts[0]
		password := parts[1]

		return m.authService.ValidateCredentials(username, password)
	}

	return false
}

// RequireAuth is an HTTP handler wrapper that requires authentication
func (m *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !m.CheckAuth(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
