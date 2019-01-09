package project

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/model"
)

// GetFeaturesForID - get project features for project id
func GetFeaturesForID(db *db.DB, id string) ([]model.Feature, *err.Error) {
	var features []model.Feature

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT f.id_feature,
									  f.name,
									  f.description 
									  FROM Project p 
									  INNER JOIN Project_feature pf ON p.id_project = pf.id_project
									  INNER JOIN Feature f ON f.id_feature = pf.id_feature
									  WHERE p.id_project = ?;`, id)

		for rows.Next() {
			var f model.Feature
			err = rows.Scan(&f.ID,
				&f.Name,
				&f.Description)

			if err != nil {
				return db.HandleError(err)
			}

			features = append(features, f)
		}

		return nil
	})

	return features, err
}
