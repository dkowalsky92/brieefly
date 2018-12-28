package project

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/model"
	"github.com/brieefly/model/project"
)

// GetDetailsForURL - get project details for project url
func GetDetailsForURL(db *db.DB, url string) (*project.Details, error) {
	var details *project.Details

	err := db.WithTransaction(func(tx *sql.Tx) error {
		projectRow := tx.QueryRow(`SELECT p.id_project,
										  p.name, 
										  p.type, 
										  p.description, 
										  p.date_created, 
										  p.date_deadline,
										  cp.name,
										  cp.image_url,
										  s.id_status,
										  s.name,
										  c.id_cms,
										  c.name,
										  c.description,
										  (SELECT AVG(op.grade) FROM Opinion op WHERE op.id_project = p.id_project GROUP BY op.id_opinion) as "avgOpn"
										  FROM Project p
										  INNER JOIN Offer o ON o.id_project = p.id_project
										  INNER JOIN Agency a ON a.id_company = o.id_company
										  INNER JOIN Company cp ON a.id_company = cp.id_company
										  INNER JOIN Status s ON s.id_status = p.id_status
										  INNER JOIN Cms c ON c.id_cms = p.id_cms
										  WHERE p.url_name = ? AND o.is_chosen = true`, url)

		var s model.ProjectStatus
		var c model.CMS
		var d project.Details
		var avgOp float64

		err := projectRow.Scan(&d.ProjectID,
			&d.Name,
			&d.Type,
			&d.Description,
			&d.DateCreated,
			&d.DateDeadline,
			&d.AgencyName,
			&d.AgencyLogoURL,
			&s.ID,
			&s.Name,
			&c.ID,
			&c.Name,
			&c.Description,
			&avgOp)

		if err != nil {
			return err
		}

		f, err := GetFeaturesForID(db, d.ProjectID)
		vi, err := GetVisualIdentitiesForID(db, d.ProjectID)
		cl, err := GetColorsForID(db, d.ProjectID)
		cf, err := GetCustomFeaturesForID(db, d.ProjectID)
		sp, err := GetSimilarProjectsForID(db, d.ProjectID)
		tg, err := GetTargetGroupsForID(db, d.ProjectID)

		d.Cms = &c
		d.Status = &s
		d.Features = f
		d.CustomFeatures = cf
		d.VisualIdentities = vi
		d.Colors = cl
		d.TargetGroups = tg
		d.SimilarProjects = sp
		d.AverageOpinion = avgOp

		details = &d

		return err
	})

	return details, err
}
