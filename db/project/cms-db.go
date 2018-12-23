package project

import (
	"database/sql"
	"fmt"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/model"
)

// GetCMSForID - gets the project's name for project id
func GetCMSForID(_db *db.DB, id string) (*model.CMS, error) {
	var cms model.CMS
	var err error

	_db.WithTransaction(func(tx *sql.Tx) error {
		row := tx.QueryRow(`SELECT c.id_cms, c.name, c.description FROM Project p
							RIGHT JOIN Cms c ON p.id_cms = c.id_cms
							WHERE p.id_project = ?;`, id)
		err = row.Scan(&cms.ID,
			&cms.Name,
			&cms.Description)

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

	return &cms, err
}
