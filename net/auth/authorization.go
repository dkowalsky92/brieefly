package auth

import (
	"net/http"
	"strings"

	"github.com/brieefly/config"
	_err "github.com/brieefly/err"
	jwt "github.com/dgrijalva/jwt-go"
)

// ValidateTokenMiddleware - validate if token is present
func ValidateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		token, pErr := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			config := config.FromContext(r.Context())
			return PublicKey(config)
		})
		if pErr == nil {
			if token.Valid {
				next.ServeHTTP(w, r)
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
