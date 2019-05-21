package agency

import (
	"database/sql"

	"github.com/brieefly/server/ctrl/agency/body"
	"github.com/brieefly/server/db"
	"github.com/brieefly/server/err"
	"github.com/brieefly/server/model"
)

// DbGetForURL - get agency for url
func DbGetForURL(db *db.DB, url string) (*model.Agency, *err.Error) {
	var agency *model.Agency

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT a.agency_code,
									a.nip_number, 
									c.id_company, 
									c.email,
									c.name, 
									c.phone, 
									c.address, 
									c.website_url, 
									c.image_url,
									c.url_name,  
									c.description, 
									c.date_last_modified, 
									c.date_created FROM Agency a
									INNER JOIN Company c ON a.id_company = c.id_company
									WHERE c.url_name = ?`, url)

		var c model.Company
		var a model.Agency

		err := row.Scan(&a.AgencyCode,
			&a.NipNumber,
			&c.ID,
			&c.Email,
			&c.Name,
			&c.Phone,
			&c.Address,
			&c.WebsiteURL,
			&c.ImageURL,
			&c.NameURL,
			&c.Description,
			&c.DateLastModified,
			&c.DateCreated)

		if err != nil {
			return db.HandleError(err)
		}

		a.Company = c

		agency = &a

		return nil
	})

	return agency, err
}

// DbGetForID - get agency for id
func DbGetForID(db *db.DB, id string) (*model.Agency, *err.Error) {
	var agency *model.Agency

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT a.agency_code,
									a.nip_number, 
									c.id_company, 
									c.email,
									c.name, 
									c.phone, 
									c.address, 
									c.website_url, 
									c.image_url, 
									c.description, 
									c.date_last_modified, 
									c.date_created FROM Agency a
									INNER JOIN Company c ON a.id_company = c.id_company
									WHERE c.id_company = ?`, id)

		var c model.Company
		var a model.Agency

		err := row.Scan(&a.AgencyCode,
			&a.NipNumber,
			&c.ID,
			&c.Email,
			&c.Name,
			&c.Phone,
			&c.Address,
			&c.WebsiteURL,
			&c.ImageURL,
			&c.Description,
			&c.DateLastModified,
			&c.DateCreated)

		if err != nil {
			return db.HandleError(err)
		}

		a.Company = c

		agency = &a

		return nil
	})

	return agency, err
}

// DbGetAll - Get all agencies
func DbGetAll(_db *db.DB) ([]body.AgencyDetails, *err.Error) {
	agencies := []body.AgencyDetails{}

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT a.agency_code,
								a.nip_number, 
								c.id_company, 
								c.email,
								c.name, 
								c.phone, 
								c.address, 
								c.website_url, 
								c.image_url, 
								c.description, 
								c.date_last_modified, 
								c.date_created
								FROM Agency a 
								INNER JOIN Company c ON a.id_company = c.id_company`)

		if err != nil {
			return _db.HandleError(err)
		}

		for rows.Next() {
			var d body.AgencyDetails
			var a model.Agency
			var c model.Company

			err = rows.Scan(&a.AgencyCode,
				&a.NipNumber,
				&c.ID,
				&c.Email,
				&c.Name,
				&c.Phone,
				&c.Address,
				&c.WebsiteURL,
				&c.ImageURL,
				&c.Description,
				&c.DateLastModified,
				&c.DateCreated)
			if err != nil {
				return _db.HandleError(err)
			}

			ao, err := DbGetAverageOpinionForAgencyID(_db, c.ID)
			if err != nil {
				return err
			}

			a.Company = c
			d.Agency = &a
			d.AverageOpinion = db.NullFloat64{Float64: *ao, Valid: true}

			agencies = append(agencies, d)
		}

		return nil
	})

	return agencies, err
}
