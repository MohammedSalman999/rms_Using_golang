package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/MohammedSalman999/rms.go/utils"
)

type contextKey string

const (
	ContextKeyEmail = contextKey("email")
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		email, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyEmail, email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
