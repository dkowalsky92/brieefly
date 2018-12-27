package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
	"github.com/brieefly/model/project"
)

// GetAllForUserID - Get all projects for user id
func GetAllForUserID(db *db.DB, id string) ([]project.UserProject, error) {
	projects := []project.UserProject{}

	err := db.WithTransaction(func(tx *sql.Tx) error {
		rows, err := tx.Query(`SELECT p.id_project,
									   p.name,
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
									   INNER JOIN Offer o ON o.id_project = p.id_project
									   INNER JOIN Agency a ON o.id_company = a.id_company  
									   INNER JOIN Company c ON c.id_company = a.id_company
									   WHERE cp.id_user = ? AND o.is_chosen = true`, id)

		if err != nil {
			log.Error(fmt.Sprintf("Error occurred: %+v", err))
			return err
		}
		for rows.Next() {
			var up project.UserProject
			var p model.Project

			err = rows.Scan(&p.ID,
				&p.Name,
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
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			c, err := GetCMSForID(db, p.ID)
			s, err := GetStatusForID(db, p.ID)

			if err != nil {
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			p.Cms = c
			p.Status = s
			up.Project = p

			projects = append(projects, up)
		}

		return err
	})

	return projects, err
}
