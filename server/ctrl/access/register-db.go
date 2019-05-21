package access

import (
	"database/sql"

	"github.com/brieefly/server/ctrl/access/body"
	"github.com/brieefly/server/ctrl/user"
	_db "github.com/brieefly/server/db"
	"github.com/brieefly/server/err"
	"github.com/brieefly/server/util"
)

// DbRegister - registers a user
func DbRegister(db *_db.DB, email string, password string) (*body.BasicUser, *err.Error) {
	var u *body.BasicUser

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		id := user.DbExists(db, email, password)
		if id.Valid == true {
			return db.HandleTypedError(nil, _db.ErrAlreadyExists)
		}

		newID := util.UUID()
		stmt, pErr := tx.Prepare(`INSERT INTO User
								  (id_user,
								   email,
							 	   password)
								  VALUES (?, ?, ?)`)
		if pErr != nil {
			return db.HandleError(pErr)
		}

		_, eErr := stmt.Exec(newID.String(), email, password)
		if eErr != nil {
			return db.HandleError(eErr)
		}

		u = &body.BasicUser{UserID: newID.String(), Email: email}

		return nil
	})

	return u, err
}
