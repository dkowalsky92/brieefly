package project

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/model"
)

// DbGetCMSForID - gets the project's name for project id
func DbGetCMSForID(_db *db.DB, id string) (*model.CMS, *err.Error) {
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

// DbGetAllCMS - 
func DbGetAllCMS(_db *db.DB) ([]model.CMS, *err.Error) {
	var cms []model.CMS

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, sqlErr := tx.Query(`SELECT c.id_cms,
								   c.name,
								   c.description
								   FROM Cms c;`)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		for rows.Next() {
			var c model.CMS
			sqlErr = rows.Scan(&c.ID, &c.Name, &c.Description)
			if sqlErr != nil {
				return _db.HandleError(sqlErr)
			}

			cms = append(cms, c)
		}

		return nil
	})

	return cms, err
}