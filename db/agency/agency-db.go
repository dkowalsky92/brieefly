package agency

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// GetForID - get agency for id
func GetForID(db *db.DB, id string) (*model.Agency, error) {
	var agency model.Agency

	err := db.WithTransaction(func(tx *sql.Tx) error {
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
		err := row.Scan(&agency.AgencyCode,
			&agency.NipNumber,
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
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		agency.Company = c

		return err
	})

	return &agency, err
}

// GetAll - Get all agencies
func GetAll(db *db.DB) ([]model.Agency, error) {
	agencies := []model.Agency{}

	err := db.WithTransaction(func(tx *sql.Tx) error {
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
			log.Error(fmt.Sprintf("Error occurred: %+v", err))
			return err
		}
		for rows.Next() {
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
				switch err {
				case sql.ErrNoRows:
					log.Error(fmt.Sprintf("No rows found"))
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			a.Company = c

			agencies = append(agencies, a)
		}

		return err
	})

	return agencies, err
}
