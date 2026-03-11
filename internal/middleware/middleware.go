package middleware

import (
	"GymMembership-api/internal/service"
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}
		
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		claims, err := service.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, `{"error":"invalid token"}`, http.StatusUnauthorized)
			return
		}

		userId := int64(claims["user_id"].(float64))
		role := claims["role"].(string)

		ctx := context.WithValue(r.Context(), "user_id", userId)
		ctx = context.WithValue(ctx, "role", role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
