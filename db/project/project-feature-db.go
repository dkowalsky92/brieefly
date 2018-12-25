package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// GetFeaturesForID - get project features for project id
func GetFeaturesForID(db *db.DB, id string) ([]model.Feature, error) {
	var features []model.Feature
	var err error

	db.WithTransaction(func(tx *sql.Tx) error {
		rows, err := tx.Query(`SELECT f.id_feature, f.name, f.description FROM Project p 
							LEFT JOIN Project_feature pf ON p.id_project = pf.id_project
							LEFT JOIN Feature f ON f.id_feature = pf.id_feature WHERE p.id_project = ?;`, id)

		for rows.Next() {
			var f model.Feature
			err = rows.Scan(&f.ID, &f.Name, &f.Description)

			if err != nil {
				switch err {
				case sql.ErrNoRows:
					log.Error(fmt.Sprintf("No rows found"))
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			features = append(features, f)
		}

		return nil
	})

	return features, err
}
