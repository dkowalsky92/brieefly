package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// GetCMSForID - gets the project's name for project id
func GetCMSForID(_db *db.DB, id string) (*model.CMS, error) {
	var cms *model.CMS

	err := _db.WithTransaction(func(tx *sql.Tx) error {
		row := tx.QueryRow(`SELECT c.id_cms,
								   c.name,
								   c.description
								   FROM Project p
							       INNER JOIN Cms c ON p.id_cms = c.id_cms
								   WHERE p.id_project = ?;`, id)
		var c model.CMS

		err := row.Scan(&c.ID,
			&c.Name,
			&c.Description)

		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		cms = &c

		return err
	})

	return cms, err
}
