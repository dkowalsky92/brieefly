package market

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model/market"
)

// GetPendingProjects - get all offers
func GetPendingProjects(db *db.DB) ([]market.PendingProject, error) {
	var projects []market.PendingProject

	err := db.WithTransaction(func(tx *sql.Tx) error {
		rows, err := tx.Query(`SELECT p.id_project,
									  p.name, 
									  p.type, 
									  p.description, 
									  p.language, 
									  p.budget_min, 
									  p.budget_max,
									  p.date_created, 
									  p.date_deadline, 
									  COUNT(cf.id_project) as "customFeatureCount",
									  GROUP_CONCAT(DISTINCT vi.type) as "visualIdentityType"
									  FROM Project p
									  INNER JOIN Cms c ON p.id_cms = c.id_cms
									  INNER JOIN Custom_feature cf ON p.id_project = cf.id_project
									  INNER JOIN Visual_identity vi ON p.id_project = vi.id_project
									  INNER JOIN Status s ON p.id_status = s.id_status
									  WHERE s.name = ?
									  GROUP BY p.id_project;`, `Pending`)
		log.Info(err)
		for rows.Next() {
			var pp market.PendingProject
			err = rows.Scan(&pp.ID,
				&pp.Name,
				&pp.Type,
				&pp.Description,
				&pp.Language,
				&pp.BudgetMin,
				&pp.BudgetMax,
				&pp.DateCreated,
				&pp.DateDeadline,
				&pp.CustomFeatureCount,
				&pp.VisualIdentityType)

			if err != nil {
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			projects = append(projects, pp)
		}

		return err
	})

	return projects, err
}
