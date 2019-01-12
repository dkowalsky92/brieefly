package login

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/brieefly/config"
	_db "github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/net/auth"
	jwt "github.com/dgrijalva/jwt-go"
)

// AuthToken - authorization token
type AuthToken struct {
	Token string `json:"token"`
}

// Info - user login info
type Info struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// JWTClaims - claims to be used inside JWT token
type JWTClaims struct {
	UserID string `json:"id_user"`
	jwt.StandardClaims
}

// Login - Login user
func Login(db *_db.DB, cnf *config.Config, email, password string) (*AuthToken, *err.Error) {
	var authToken *AuthToken

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT u.id_user FROM User u
							WHERE u.email = ? AND u.password = ?;`, email, password)

		var id string

		err := row.Scan(&id)

		if err != nil {
			return db.HandleTypedError(err, _db.ErrUserNotFound)
		}

		claims := &JWTClaims{
			id,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 15).Unix(),
				Issuer:    "Brieefly",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

		pkey, err := auth.PrivateKey(cnf)
		if err != nil {
			return db.HandleError(err)
		}

		ss, err := token.SignedString(pkey)
		fmt.Println(ss)
		if err != nil {
			return db.HandleError(err)
		}

		authToken = &AuthToken{ss}

		return db.HandleError(err)
	})

	return authToken, err
}
