package market

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/model/market"
)

// GetPendingProjects - get all offers
func GetPendingProjects(db *db.DB) ([]market.PendingProject, *err.Error) {
	var projects []market.PendingProject

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
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
				return db.HandleError(err)
			}

			projects = append(projects, pp)
		}

		return nil
	})

	return projects, err
}
