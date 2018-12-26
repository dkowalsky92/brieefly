package market

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// GetOffersForID - get offers for project id
func GetOffersForID(db *db.DB, id string) ([]model.Offer, error) {
	var offers []model.Offer

	err := db.WithTransaction(func(tx *sql.Tx) error {
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
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			offers = append(offers, o)
		}

		return err
	})

	return offers, err
}

// GetAllOffers - get all offers
func GetAllOffers(db *db.DB) ([]model.Offer, error) {
	var offers []model.Offer

	err := db.WithTransaction(func(tx *sql.Tx) error {
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
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			offers = append(offers, o)
		}

		return err
	})

	return offers, err
}
