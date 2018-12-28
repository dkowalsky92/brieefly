package agency

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
	"github.com/brieefly/model/agency"
)

// GetDetailsForName - get agency details for agency name
func GetDetailsForName(db *db.DB, name string) (agency.Details, error) {
	var details agency.Details

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
									  a.id_company 
									  FROM Project p
									  INNER JOIN Offer o ON o.id_project = p.id_project
									  INNER JOIN Agency a ON a.id_company = o.id_company
									  INNER JOIN Company c ON c.id_company = a.id_company
									  INNER JOIN Status s ON s.id_status = p.id_status
									  WHERE o.is_chosen = true AND s.name = ? AND c.name LIKE ?`, `Finished`, name)
		if err != nil {
			log.Error(fmt.Sprintf("Error occurred: %+v", err))
			return err
		}

		p := []model.Project{}
		var agnID string

		for rows.Next() {
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
				&agnID)

			if err != nil {
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}
		}

		a, err := GetForID(db, agnID)

		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		details.Agency = a
		details.FinishedProjects = p

		return err
	})

	return details, err
}
