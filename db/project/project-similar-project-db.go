package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// GetSimilarProjectsForID - get similar projects for project id
func GetSimilarProjectsForID(db *db.DB, id string) ([]model.SimilarProject, error) {
	var similarProjects []model.SimilarProject

	err := db.WithTransaction(func(tx *sql.Tx) error {
		rows, err := tx.Query(`SELECT sp.id_similar_project,
									  sp.project_url,
									  sp.id_project
									  FROM Similar_project sp
									  INNER JOIN Project p ON p.id_project = sp.id_project
									  WHERE p.id_project = ?;`, id)
		fmt.Println(err)
		for rows.Next() {
			var sp model.SimilarProject
			err = rows.Scan(&sp.ID,
				&sp.ProjectURL,
				&sp.ProjectID)

			if err != nil {
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			similarProjects = append(similarProjects, sp)
		}

		return err
	})

	return similarProjects, err
}
