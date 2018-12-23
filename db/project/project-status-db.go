package project

import (
	"database/sql"
	"fmt"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/model"
)

// GetStatusForID - get status for project id
func GetStatusForID(db *db.DB, id string) (model.ProjectStatus, error) {
	var status model.ProjectStatus
	var err error

	db.WithTransaction(func(tx *sql.Tx) error {
		row := tx.QueryRow(`SELECT s.id_status, p.name FROM Project p
							RIGHT JOIN Status s ON p.id_status = s.id_status
							WHERE p.id_project = ?;`, id)

		err = row.Scan(&status.ID, &status.Name)

		if err != nil {
			switch err {
			case sql.ErrNoRows:
				log.Error(fmt.Sprintf("No rows found"))
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		return nil
	})

	return status, err
}
