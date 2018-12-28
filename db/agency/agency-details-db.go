package agency

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model/agency"
)

// GetFinishedProjectsForURL - get finished projects for company id
func GetFinishedProjectsForURL(db *db.DB, url string) ([]agency.BasicProject, error) {
	projects := []agency.BasicProject{}

	err := db.WithTransaction(func(tx *sql.Tx) error {
		rows, err := tx.Query(`SELECT p.id_project,
									  p.name,
									  p.type,
									  p.description,
									  p.image_url,
									  (SELECT AVG(oo.grade) FROM Opinion oo WHERE oo.id_project = p.id_project) as "avgOpn"
									  FROM Project p
									  INNER JOIN Offer o ON o.id_project = p.id_project
									  INNER JOIN Agency a ON a.id_company = o.id_company
									  INNER JOIN Company c ON c.id_company = a.id_company
									  INNER JOIN Status s ON s.id_status = p.id_status
									  WHERE o.is_chosen = true AND s.name = ? AND c.url_name = ?`, `Finished`, url)
		if err != nil {
			log.Error(fmt.Sprintf("Error occurred: %+v", err))
			return err
		}

		for rows.Next() {
			var bp agency.BasicProject

			err = rows.Scan(&bp.ID,
				&bp.Name,
				&bp.Type,
				&bp.Description,
				&bp.ImageURL,
				&bp.AverageOpinion)

			if err != nil {
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			projects = append(projects, bp)
		}

		return err
	})

	return projects, err
}

// GetAgencyAndOpinionsForURL - get agency details for company url
func GetAgencyAndOpinionsForURL(db *db.DB, url string) (*agency.Details, error) {
	var details *agency.Details

	err := db.WithTransaction(func(tx *sql.Tx) error {
		a, err := GetForURL(db, url)

		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		var d agency.Details

		d.Agency = a

		row := tx.QueryRow(`SELECT AVG(op.grade)
						 	FROM Opinion op
							INNER JOIN Offer o ON o.id_project = op.id_project
							WHERE o.id_company = ?`, a.Company.ID)

		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		err = row.Scan(&d.AverageOpinion)

		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		details = &d

		return err
	})

	return details, err
}
