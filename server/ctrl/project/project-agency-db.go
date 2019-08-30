package project

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/ctrl/project/body"
)

// DbGetBiddingAgencyForURL - 
func DbGetBiddingAgencyForURL(_db *db.DB, url string) (*body.ProjectBiddingAgency, *err.Error) {
	var pba *body.ProjectBiddingAgency

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT c.id_company,
								   c.name,
								   c.image_url
								   FROM Company c
								   INNER JOIN Agency a ON c.id_company = c.id_company
								   INNER JOIN Offer o ON o.id_company = a.id_company
								   INNER JOIN Project p ON o.id_project = o.id_project
								   WHERE p.url_name = ? AND o.is_chosen = true;`, url)

		var p body.ProjectBiddingAgency

		err := row.Scan(
			&p.ID,
			&p.Name,
			&p.ImageURL)

		if err != nil {
			return _db.HandleError(err)
		}

		pba = &p

		return nil
	})

	return pba, err
}