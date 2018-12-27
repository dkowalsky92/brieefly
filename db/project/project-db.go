package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// GetAllForUserID - Get all projects for user id
func GetAllForUserID(db *db.DB, id string) ([]model.Project, error) {
	projects := []model.Project{}

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
			var p model.Project
			var agnName string

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
				&agnName)

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

			projects = append(projects, p)
		}

		return err
	})

	return projects, err
}

// GetNameForID - gets the project's name for project id
func GetNameForID(_db *db.DB, id string) (*db.NullString, error) {
	var name db.NullString
	var err error

	_db.WithTransaction(func(tx *sql.Tx) error {
		row := tx.QueryRow(`SELECT name FROM Project WHERE id_project = ?;`, id)
		err = row.Scan(&name)

		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		return nil
	})

	return &name, err
}
