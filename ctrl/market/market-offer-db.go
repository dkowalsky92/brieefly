package market

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/model"
)

// DbGetOffersForID - get offers for project id
func DbGetOffersForID(db *db.DB, id string) ([]model.Offer, *err.Error) {
	var offers []model.Offer

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT o.id_offer,
									  o.salary, 
									  o.is_chosen,
									  o.date_deadline,
									  o.date_created,
									  o.id_project,
									  o.id_company
									  FROM Offer o
									  WHERE o.id_project = ?;`, id)

		for rows.Next() {
			var o model.Offer
			err = rows.Scan(&o.ID,
				&o.Salary,
				&o.IsChosen,
				&o.DateDeadline,
				&o.DateCreated,
				&o.ProjectID,
				&o.CompanyID)

			if err != nil {
				return db.HandleError(err)
			}

			offers = append(offers, o)
		}

		return nil
	})

	return offers, err
}

// DbGetAllOffers - get all offers
func DbGetAllOffers(db *db.DB) ([]model.Offer, *err.Error) {
	var offers []model.Offer

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT o.id_offer,
									  o.salary, 
									  o.is_chosen,
									  o.date_deadline,
									  o.date_created,
									  o.id_project,
									  o.id_company
									  FROM Offer o;`)

		for rows.Next() {
			var o model.Offer
			err = rows.Scan(&o.ID,
				&o.Salary,
				&o.IsChosen,
				&o.DateDeadline,
				&o.DateCreated,
				&o.ProjectID,
				&o.CompanyID)

			if err != nil {
				return db.HandleError(err)
			}

			offers = append(offers, o)
		}

		return nil
	})

	return offers, err
}
