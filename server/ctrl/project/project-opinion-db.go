package project

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/model"
)

// DbGetAverageOpinionForID - 
func DbGetAverageOpinionForID(_db *db.DB, projectid string) (*float64, *err.Error) {
	var result *float64

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT AVG(o.grade)
									  FROM Opinion o
									  INNER JOIN Project p ON p.id_project = o.id_project
									  WHERE o.id_project = ?`, projectid)
		var avg float64
		sqlErr := row.Scan(&avg)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}		

		result = &avg

		return nil
	})

	return result, err
}

// DbGetOpinionsForID - 
func DbGetOpinionsForID(_db *db.DB, projectid string) ([]model.Opinion, *err.Error) {
	var op []model.Opinion

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, sqlErr := tx.Query(`SELECT o.id_opinion,
									  o.grade,
									  o.description,
									  o.date_created
									  FROM Opinion o
									  INNER JOIN Project p ON p.id_project = o.id_project
									  WHERE o.id_project = ?
									  ORDER BY o.date_created DESC;`, projectid)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}		

		for rows.Next() {
			var o model.Opinion
			sqlErr = rows.Scan(&o.ID,
				&o.Grade,
				&o.Description,
				&o.DateCreated)

			if sqlErr != nil {
				return _db.HandleError(sqlErr)
			}

			op = append(op, o)
		}

		return nil
	})

	return op, err
}