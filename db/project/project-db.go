package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
)

// GetNameForID - gets the project's name for project id
func GetNameForID(_db *db.DB, id string) (*db.NullString, error) {
	var name db.NullString

	err := _db.WithTransaction(func(tx *sql.Tx) error {
		row := tx.QueryRow(`SELECT name FROM Project WHERE id_project = ?;`, id)
		err := row.Scan(&name)

		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		return err
	})

	return &name, err
}
