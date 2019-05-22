package project

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/model"
)

// DbGetCustomFeaturesForID - get project's custom features for project id
func DbGetCustomFeaturesForID(db *db.DB, id string) ([]model.CustomFeature, *err.Error) {
	var customFeatures []model.CustomFeature

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT cf.id_custom_feature,
									  cf.name,
									  cf.description
									  FROM Custom_feature cf 
									  INNER JOIN Project p ON p.id_project = cf.id_project
									  WHERE p.id_project = ?;`, id)

		for rows.Next() {
			var cf model.CustomFeature
			err = rows.Scan(&cf.ID,
				&cf.Name,
				&cf.Description)

			if err != nil {
				return db.HandleError(err)
			}

			customFeatures = append(customFeatures, cf)
		}

		return nil
	})

	return customFeatures, err
}
