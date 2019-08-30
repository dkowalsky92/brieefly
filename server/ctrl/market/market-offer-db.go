package market

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/ctrl/market/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/model"
)

// DbGetOffersForID - get offers for project id
func DbGetOffersForID(db *db.DB, id string) ([]model.Offer, *err.Error) {
	var offers []model.Offer

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT o.id_offer,
									  o.salary_min,
									  o.salary_max, 
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
				&o.SalaryMin,
				&o.SalaryMax,
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
									  o.salary_min,
									  o.salary_max,  
									  o.is_chosen,
									  o.date_deadline,
									  o.date_created,
									  o.id_project,
									  o.id_company
									  FROM Offer o;`)

		for rows.Next() {
			var o model.Offer
			err = rows.Scan(&o.ID,
				&o.SalaryMin,
				&o.SalaryMax,
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

// DbInsertOffer -
func DbInsertOffer(_db *db.DB, bd body.OfferBody, userid string) *err.Error {
	txErr := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT a.id_company FROM Agency a
							INNER JOIN Agency_employee ae ON a.id_company = ae.id_company
							WHERE ae.id_user = ?`, userid)
		var companyid string
		sqlErr := row.Scan(&companyid)
		if sqlErr != nil || companyid == "" {
			return _db.HandleError(sqlErr)
		}

		offer := body.NewOffer(bd, companyid)
		offerStmt := db.InsertStmt(tx, offer, "Offer")
		_, sqlErr = offerStmt.Stmt.Exec(offerStmt.Args...)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		return nil
	})

	return txErr
}
