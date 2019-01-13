package project

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/model"
)

// DbGetSimilarProjectsForID - get similar projects for project id
func DbGetSimilarProjectsForID(db *db.DB, id string) ([]model.SimilarProject, *err.Error) {
	var similarProjects []model.SimilarProject

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT sp.id_similar_project,
									  sp.project_url,
									  sp.id_project
									  FROM Similar_project sp
									  INNER JOIN Project p ON p.id_project = sp.id_project
									  WHERE p.id_project = ?;`, id)
		for rows.Next() {
			var sp model.SimilarProject
			err = rows.Scan(&sp.ID,
				&sp.ProjectURL,
				&sp.ProjectID)

			if err != nil {
				return db.HandleError(err)
			}

			similarProjects = append(similarProjects, sp)
		}

		return nil
	})

	return similarProjects, err
}
