package project

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/ctrl/project/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
)

// DbGetDetailsForURL - get project details for project url
func DbGetDetailsForURL(db *db.DB, url string) (*body.ProjectDetailsBundle, *err.Error) {
	var bundle *body.ProjectDetailsBundle

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		projectRow := tx.QueryRow(`SELECT p.id_project,
										  p.name, 
										  p.type, 
										  p.description, 
										  p.date_created, 
										  p.date_deadline
										  FROM Project p
										  WHERE p.url_name = ?`, url)
		var d body.ProjectDetails

		err := projectRow.Scan(&d.ProjectID,
			&d.Name,
			&d.Type,
			&d.Description,
			&d.DateCreated,
			&d.DateDeadline)

		if err != nil {
			return db.HandleError(err)
		}

		pba, _ := DbGetBiddingAgencyForURL(db, url)
		cms, _ := DbGetCMSForID(db, d.ProjectID)
		s, _ := DbGetStatusForID(db, d.ProjectID)
		avgOp, _ := DbGetAverageOpinionForID(db, d.ProjectID)

		f, fErr := DbGetFeaturesForID(db, d.ProjectID)
		if fErr != nil {
			return fErr
		}
		vi, viErr := DbGetVisualIdentitiesForID(db, d.ProjectID)
		if viErr != nil {
			return viErr
		}
		cl, cErr := DbGetColorsForID(db, d.ProjectID)
		if cErr != nil {
			return cErr
		}
		cf, cfErr := DbGetCustomFeaturesForID(db, d.ProjectID)
		if cfErr != nil {
			return cfErr
		}
		sp, spErr := DbGetSimilarProjectsForID(db, d.ProjectID)
		if spErr != nil {
			return spErr
		}
		tg, tgErr := DbGetTargetGroupsForID(db, d.ProjectID)
		if tgErr != nil {
			return tgErr
		}
		op, opErr := DbGetOpinionsForID(db, d.ProjectID)
		if opErr != nil {
			return opErr
		}
		
		if avgOp != nil {
			d.AverageOpinion = *avgOp
		} else {
			d.AverageOpinion = 0
		}

		bundle = &body.ProjectDetailsBundle{ProjectDetails: d, Bidder: pba, Status: s, CMS: cms}
		bundle.Features = f
		bundle.CustomFeatures = cf
		bundle.VisualIdentities = vi
		bundle.Colors = cl
		bundle.TargetGroups = tg
		bundle.SimilarProjects = sp
		bundle.Opinions = op

		return nil
	})

	return bundle, err
}
