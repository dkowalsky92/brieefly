package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// GetStatusForID - get status for project id
func GetStatusForID(db *db.DB, id string) (model.ProjectStatus, error) {
	var status model.ProjectStatus
	var err error

	db.WithTransaction(func(tx *sql.Tx) error {
		row := tx.QueryRow(`SELECT s.id_status,
							       p.name
							       FROM Project p
								   INNER JOIN Status s ON p.id_status = s.id_status
								   WHERE p.id_project = ?;`, id)

		err = row.Scan(&status.ID, &status.Name)

		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		return nil
	})

	return status, err
}
