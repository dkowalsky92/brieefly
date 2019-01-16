package project

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/model"
)

// DbGetStatusForID - get status for project id
func DbGetStatusForID(db *db.DB, id string) (*model.ProjectStatus, *err.Error) {
	var status *model.ProjectStatus

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT s.id_status,
							       s.name
							       FROM Project p
								   INNER JOIN Status s ON p.id_status = s.id_status
								   WHERE p.id_project = ?;`, id)
		var s model.ProjectStatus

		err := row.Scan(&s.ID,
			&s.Name)

		if err != nil {
			return db.HandleError(err)
		}

		status = &s

		return nil
	})

	return status, err
}

// DbGetStatusForName - get status for name
func DbGetStatusForName(db *db.DB, name string) (*model.ProjectStatus, *err.Error) {
	var status *model.ProjectStatus

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT s.id_status,
							       s.name
							       FROM Status s
								   WHERE s.name = ?;`, name)
		var s model.ProjectStatus

		err := row.Scan(&s.ID,
			&s.Name)

		if err != nil {
			return db.HandleError(err)
		}

		status = &s

		return nil
	})

	return status, err
}
