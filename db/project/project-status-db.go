package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// GetStatusForID - get status for project id
func GetStatusForID(db *db.DB, id string) (*model.ProjectStatus, error) {
	var status *model.ProjectStatus

	err := db.WithTransaction(func(tx *sql.Tx) error {
		row := tx.QueryRow(`SELECT s.id_status,
							       s.name
							       FROM Project p
								   INNER JOIN Status s ON p.id_status = s.id_status
								   WHERE p.id_project = ?;`, id)
		var s model.ProjectStatus

		err := row.Scan(&s.ID,
			&s.Name)

		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		status = &s

		return err
	})

	return status, err
}
