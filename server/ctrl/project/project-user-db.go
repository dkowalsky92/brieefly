package project

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/ctrl/project/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/model"
)

// DbGetAllForUserID - Get all projects for user id
func DbGetAllForUserID(db *db.DB, id string) ([]body.UserProject, *err.Error) {
	projects := []body.UserProject{}

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT p.id_project,
									   p.name,
									   p.url_name,
									   p.type,
									   p.description,
									   p.image_url,
									   p.language,
									   p.budget_min,
									   p.budget_max, 
									   p.subpage_count, 
									   p.overall_progress, 
									   p.date_created,
									   p.date_deadline,
									   p.date_last_modified,
									   c.name
									   FROM Project p
									   INNER JOIN Client_project cp ON p.id_project = cp.id_project
									   INNER JOIN Client cl ON cl.id_user = cp.id_user
									   INNER JOIN Company c ON cl.id_company = c.id_company
									   WHERE cp.id_user = ?`, id)

		if err != nil {
			return db.HandleError(err)
		}

		for rows.Next() {
			var up body.UserProject
			var p model.Project

			err = rows.Scan(&p.ID,
				&p.Name,
				&p.NameURL,
				&p.Type,
				&p.Description,
				&p.ImageURL,
				&p.Language,
				&p.BudgetMin,
				&p.BudgetMax,
				&p.SubpageCount,
				&p.OverallProgress,
				&p.DateCreated,
				&p.DateDeadline,
				&p.DateLastModified,
				&up.AgencyName)

			if err != nil {
				return db.HandleError(err)
			}

			c, _ := DbGetCMSForID(db, p.ID)
			s, sErr := DbGetStatusForID(db, p.ID)
			if sErr != nil {
				return sErr
			}

			up.CMS = c
			up.Status = s
			up.Project = p

			projects = append(projects, up)
		}

		return nil
	})

	return projects, err
}
