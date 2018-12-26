package project

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/model"
	"github.com/brieefly/model/project"
)

// GetDetailsForID - get project details for project id
func GetDetailsForID(db *db.DB, id string) (project.Details, error) {
	var details project.Details

	err := db.WithTransaction(func(tx *sql.Tx) error {
		projectRow := tx.QueryRow(`SELECT p.id_project,
										  p.name, 
										  p.type, 
									  	  p.description, 
										  p.date_created, 
										  p.date_deadline,
										  s.id_status,
										  s.name,
										  c.id_cms,
										  c.name,
										  c.description
										  FROM Project p
										  INNER JOIN Status s ON s.id_status = p.id_status
                                          INNER JOIN Cms c ON c.id_cms = p.id_cms
										  WHERE p.id_project = ?;`, id)

		var s model.ProjectStatus
		var c model.CMS

		err := projectRow.Scan(&details.ProjectID,
			&details.Name,
			&details.Type,
			&details.Description,
			&details.DateCreated,
			&details.DateDeadline,
			&s.ID,
			&s.Name,
			&c.ID,
			&c.Name,
			&c.Description)

		f, err := GetFeaturesForID(db, id)
		vi, err := GetVisualIdentitiesForID(db, id)
		cl, err := GetColorsForID(db, id)
		cf, err := GetCustomFeaturesForID(db, id)
		tg, err := GetTargetGroupsForID(db, id)

		details.Cms = &c
		details.Status = &s
		details.Features = f
		details.CustomFeatures = cf
		details.VisualIdentities = vi
		details.Colors = cl
		details.TargetGroups = tg

		return err
	})

	return details, err
}
