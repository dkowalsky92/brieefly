package project

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/dkowalsky/brieefly/ctrl/project/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
)

// DbGetOffersForSlug - get similar projects for project id
func DbGetOffersForSlug(db *db.DB, slug string) ([]body.AgencyOffer, *err.Error) {
	var offers []body.AgencyOffer

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT o.id_offer,
									  o.salary_min,
									  o.salary_max,
									  o.is_chosen,
									  o.date_start,
									  o.date_deadline,
									  o.date_created,
									  o.id_project,
									  o.id_company,
									  c.name
									  FROM Offer o
									  INNER JOIN Project p ON p.id_project = o.id_project
									  INNER JOIN Company c ON c.id_company = o.id_company
									  WHERE p.url_name = ?
									  ORDER BY o.salary_min ASC;`, slug)
		for rows.Next() {
			var o body.AgencyOffer
			err = rows.Scan(&o.ID,
				&o.SalaryMin,
				&o.SalaryMax,
				&o.IsChosen,
				&o.DateStart,
				&o.DateDeadline,
				&o.DateCreated,
				&o.ProjectID,
				&o.CompanyID,
				&o.CompanyName)

			if err != nil {
				return db.HandleError(err)
			}

			offers = append(offers, o)
		}

		return nil
	})

	return offers, err
}

// DbMarkChosen -
func DbMarkChosen(db *db.DB, idOffer, projectSlug string) *err.Error {
	var existingIDOffer *string
	txErr := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT o.id_offer FROM Offer o 
							WHERE o.is_chosen = true 
							AND o.id_project = (SELECT id_project FROM Project
												  WHERE url_name = ?)`, projectSlug)

		sqlErr := row.Scan(&existingIDOffer)
		if sqlErr != nil {
			if sqlErr == sql.ErrNoRows {
				stmt, sqlErr := tx.Prepare(`UPDATE Offer SET is_chosen = true WHERE id_offer = ?`)
				if sqlErr != nil {
					return db.HandleError(sqlErr)
				}

				_, sqlErr = stmt.Exec(idOffer)
				if sqlErr != nil {
					return db.HandleError(sqlErr)
				}

				return nil
			}

			return db.HandleError(sqlErr)
		}

		return err.New(errors.New("there is a chosen offer for this project"), http.StatusConflict, map[string]interface{}{"offerId": existingIDOffer})
	})

	return txErr
}
