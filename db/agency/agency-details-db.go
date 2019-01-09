package agency

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/model/agency"
)

// GetFinishedProjectsForURL - get finished projects for company id
func GetFinishedProjectsForURL(db *db.DB, url string) ([]agency.BasicProject, *err.Error) {
	projects := []agency.BasicProject{}

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
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
			return db.HandleError(err)
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
				return db.HandleError(err)
			}

			projects = append(projects, bp)
		}

		return nil
	})

	return projects, err
}

// GetAgencyAndOpinionsForURL - get agency details for company url
func GetAgencyAndOpinionsForURL(db *db.DB, url string) (*agency.Details, *err.Error) {
	var details *agency.Details

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		a, err := GetForURL(db, url)
		if err != nil {
			return err
		}

		var d agency.Details
		d.Agency = a

		row := tx.QueryRow(`SELECT AVG(op.grade)
						 	FROM Opinion op
							INNER JOIN Offer o ON o.id_project = op.id_project
							WHERE o.id_company = ?`, a.Company.ID)
		if err != nil {
			return err
		}

		sErr := row.Scan(&d.AverageOpinion)
		if sErr != nil {
			return db.HandleError(sErr)
		}

		details = &d

		return nil
	})

	return details, err
}
