package project

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/model"
	"github.com/brieefly/model/project"
)

// type Offer struct {
// 	ID           string        `json:"idOffer"`
// 	Salary       int64         `json:"salary"`
// 	IsChosen     bool          `json:"isChosen"`
// 	DateDeadline db.NullTime   `json:"dateDeadline"`
// 	DateCreated  db.NullTime   `json:"dateCreated"`
// 	ProjectID    db.NullString `json:"idProject"`
// 	CompanyID    string        `json:"idCompany"`
// }

// GetDetailsForID - get project details for project id
func GetDetailsForID(db *db.DB, id string) (*project.Details, error) {
	var details project.Details

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
										  WHERE p.id_project = ? AND o.is_chosen = true`, id)

		var s model.ProjectStatus
		var c model.CMS
		var avgOp float64

		err := projectRow.Scan(&details.ProjectID,
			&details.Name,
			&details.Type,
			&details.Description,
			&details.DateCreated,
			&details.DateDeadline,
			&details.AgencyName,
			&details.AgencyLogoURL,
			&s.ID,
			&s.Name,
			&c.ID,
			&c.Name,
			&c.Description,
			&avgOp)

		f, err := GetFeaturesForID(db, id)
		vi, err := GetVisualIdentitiesForID(db, id)
		cl, err := GetColorsForID(db, id)
		cf, err := GetCustomFeaturesForID(db, id)
		sp, err := GetSimilarProjectsForID(db, id)
		tg, err := GetTargetGroupsForID(db, id)

		details.Cms = &c
		details.Status = &s
		details.Features = f
		details.CustomFeatures = cf
		details.VisualIdentities = vi
		details.Colors = cl
		details.TargetGroups = tg
		details.SimilarProjects = sp
		details.AverageOpinion = avgOp

		return err
	})

	return &details, err
}
