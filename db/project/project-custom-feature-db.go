package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// GetCustomFeaturesForID - get project's custom features for project id
func GetCustomFeaturesForID(db *db.DB, id string) ([]model.CustomFeature, error) {
	var customFeatures []model.CustomFeature

	err := db.WithTransaction(func(tx *sql.Tx) error {
		rows, err := tx.Query(`SELECT cf.id_custom_feature,
									  cf.name,
									  cf.description, 
									  cf.id_project
									  FROM Custom_feature cf 
									  INNER JOIN Project p ON p.id_project = cf.id_project
									  WHERE p.id_project = ?;`, id)

		for rows.Next() {
			var cf model.CustomFeature
			err = rows.Scan(&cf.ID,
				&cf.Name,
				&cf.Description,
				&cf.ProjectID)

			if err != nil {
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			customFeatures = append(customFeatures, cf)
		}

		return err
	})

	return customFeatures, err
}
