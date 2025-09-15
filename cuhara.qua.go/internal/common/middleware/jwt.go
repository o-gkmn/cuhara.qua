package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const (
	CtxUserIDKey   ctxKey = "userID"
	CtxEmailKey    ctxKey = "email"
	CtxRoleIDKey   ctxKey = "roleID"
	CtxTenantIDKey ctxKey = "tennantID"
)

func JWT(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if strings.HasPrefix(auth, "Bearer ") {
				http.Error(w, "missing bearer token", http.StatusUnauthorized)
				return
			}
			tokenStr := strings.TrimPrefix(auth, "Bearer ")

			tok, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrTokenSignatureInvalid
				}
				return []byte(secret), nil
			})
			if err != nil || !tok.Valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
			}

			claims, ok := tok.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "invalid claims", http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			if v, ok := claims["sub"].(string); ok {
				ctx = context.WithValue(ctx, CtxUserIDKey, v)
			}
			if v, ok := claims["email"].(string); ok {
				ctx = context.WithValue(ctx, CtxEmailKey, v)
			}
			if v, ok := claims["roleId"].(float64); ok {
				ctx = context.WithValue(ctx, CtxRoleIDKey, int64(v))
			}
			if v, ok := claims["tennantId"].(float64); ok {
				ctx = context.WithValue(ctx, CtxTenantIDKey, int64(v))
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
