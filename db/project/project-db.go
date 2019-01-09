package project

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/err"
)

// GetNameForID - gets the project's name for project id
func GetNameForID(_db *db.DB, id string) (*db.NullString, *err.Error) {
	var name db.NullString

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT name FROM Project WHERE id_project = ?;`, id)
		err := row.Scan(&name)

		if err != nil {
			return _db.HandleError(err)
		}

		return nil
	})

	return &name, err
}
