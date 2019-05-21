package project

import (
	"database/sql"

	"github.com/brieefly/server/ctrl/project/body"
	"github.com/brieefly/server/db"
	"github.com/brieefly/server/err"
)

// -- Table: Offer
// CREATE TABLE Offer (
//     id_offer int NOT NULL AUTO_INCREMENT,
//     salary int NOT NULL,
//     is_chosen bool NOT NULL DEFAULT false,
//     date_deadline date NOT NULL,
//     date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     id_project int NULL,
//     id_company int NOT NULL,
//     CONSTRAINT Offer_pk PRIMARY KEY (id_offer)
// );

// DbGetOffersForSlug - get similar projects for project id
func DbGetOffersForSlug(db *db.DB, slug string) ([]body.AgencyOffer, *err.Error) {
	var offers []body.AgencyOffer

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT o.id_offer,
									  o.salary,
									  o.is_chosen,
									  o.date_deadline,
									  o.date_created,
									  o.id_project,
									  o.id_company,
									  c.name
									  FROM Offer o
									  INNER JOIN Project p ON p.id_project = o.id_project
									  INNER JOIN Company c ON c.id_company = o.id_company
									  WHERE p.url_name = ?
									  ORDER BY o.salary DESC;`, slug)
		for rows.Next() {
			var o body.AgencyOffer
			err = rows.Scan(&o.ID,
				&o.Salary,
				&o.IsChosen,
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
