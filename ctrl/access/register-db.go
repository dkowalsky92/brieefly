package access

import (
	"database/sql"

	"github.com/brieefly/ctrl/access/body"
	"github.com/brieefly/ctrl/user"
	_db "github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/util"
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
