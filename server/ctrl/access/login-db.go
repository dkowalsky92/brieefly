package access

import (
	"database/sql"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/ctrl/access/body"
	"github.com/dkowalsky/brieefly/ctrl/user"
	_db "github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/net/auth"
)

// DbLogin - Login user
func DbLogin(db *_db.DB, cnf *config.Config, email, password string) (*body.UserInfo, *err.Error) {
	var info *body.UserInfo

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		ui, err := DbGetUserInfo(db, cnf, email, password)
		if err != nil {
			return err
		}

		info = ui

		return nil
	})

	return info, err
}

// DbGetUserInfo -
func DbGetUserInfo(db *_db.DB, cnf *config.Config, email, password string) (*body.UserInfo, *err.Error) {
	var info *body.UserInfo

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		id := user.DbExists(db, email, password)
		if id.Valid == false {
			return db.HandleTypedError(nil, _db.ErrNotFound)
		}

		row := db.QueryRow(`SELECT u.id_user,
								   ar.role 
								   FROM User u
								   INNER JOIN Agency_employee ae ON ae.id_user = u.id_user
								   INNER JOIN Agency_role ar ON ae.id_agency_role = ar.id_agency_role
								   WHERE u.id_user = ?`, id.String)
		var ui body.UserInfo
		err := row.Scan(&ui.UserID, &ui.Role)
		if err != nil {
			return db.HandleError(err)
		}

		info = &ui

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

		info.Token = ss

		return nil
	})

	return info, err
}
