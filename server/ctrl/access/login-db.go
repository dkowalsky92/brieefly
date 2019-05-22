package access

import (
	"database/sql"
	"time"

	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/ctrl/access/body"
	"github.com/dkowalsky/brieefly/ctrl/user"
	_db "github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/net/auth"
	jwt "github.com/dgrijalva/jwt-go"
)

// DbLogin - Login user
func DbLogin(db *_db.DB, cnf *config.Config, email, password string) (*body.Token, *err.Error) {
	var authToken *body.Token

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		id := user.DbExists(db, email, password)
		if id.Valid == false {
			return db.HandleTypedError(nil, _db.ErrNotFound)
		}

		claims := &auth.JWTClaims{
			id.String,
			jwt.StandardClaims{
				IssuedAt:  time.Now().Unix(),
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
		if err != nil {
			return db.HandleError(err)
		}

		authToken = &body.Token{Token: ss, UserID: id.String}

		return db.HandleError(err)
	})

	return authToken, err
}
