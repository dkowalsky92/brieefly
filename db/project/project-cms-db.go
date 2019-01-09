package project

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/model"
)

// GetCMSForID - gets the project's name for project id
func GetCMSForID(_db *db.DB, id string) (*model.CMS, *err.Error) {
	var cms *model.CMS

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
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
			return _db.HandleError(err)
		}

		cms = &c

		return nil
	})

	return cms, err
}
