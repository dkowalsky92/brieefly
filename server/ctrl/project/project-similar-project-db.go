package project

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/model"
)

// DbGetSimilarProjectsForID - get similar projects for project id
func DbGetSimilarProjectsForID(db *db.DB, id string) ([]model.SimilarProject, *err.Error) {
	var similarProjects []model.SimilarProject

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT sp.id_similar_project,
									  sp.project_url
									  FROM Similar_project sp
									  INNER JOIN Project p ON p.id_project = sp.id_project
									  WHERE p.id_project = ?;`, id)
		for rows.Next() {
			var sp model.SimilarProject
			err = rows.Scan(&sp.ID,
				&sp.ProjectURL)

			if err != nil {
				return db.HandleError(err)
			}

			similarProjects = append(similarProjects, sp)
		}

		return nil
	})

	return similarProjects, err
}
