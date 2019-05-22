package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/dkowalsky/brieefly/config"
	_err "github.com/dkowalsky/brieefly/err"
	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const (
	userIDKey contextKey = "userIDKey"
)

// JWTClaims - claims to be used inside JWT token
type JWTClaims struct {
	UserID string `json:"idUser"`
	jwt.StandardClaims
}

// ValidateTokenMiddleware - validate if token is present
func ValidateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if h == "" {
			err := _err.New(errors.New("invalid token, access restricted"), http.StatusUnauthorized, map[string]interface{}{})
			_err.WriteError(err, w)
			return
		}
		value := strings.TrimPrefix(h, "Bearer ")
		token, pErr := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
			config := config.FromContext(r.Context())
			return PublicKey(config)
		})
		if pErr == nil {
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				ctx := context.WithValue(r.Context(), userIDKey, claims["idUser"])
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				err := _err.New(pErr, http.StatusUnauthorized, map[string]interface{}{})
				_err.WriteError(err, w)
			}
		} else {
			err := _err.New(pErr, http.StatusUnauthorized, map[string]interface{}{})
			_err.WriteError(err, w)
		}
	})
}

// UserIDFromContext -
func UserIDFromContext(ctx context.Context) *string {
	val := ctx.Value(userIDKey).(string)
	if val == "" {
		return nil
	}
	return &val
}
