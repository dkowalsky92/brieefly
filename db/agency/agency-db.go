package agency

import (
	"database/sql"
	"fmt"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/model"
)

// GetAll - Get all agencies
func GetAll(db *db.DB) ([]model.Agency, error) {
	agencies := []model.Agency{}
	var err error

	db.WithTransaction(func(tx *sql.Tx) error {
		rows, qerr := tx.Query("SELECT a.agency_code, a.nip_number, c.id_company, c.email, c.phone, c.address, c.website_url, c.image_url, c.description, c.date_last_modified, c.date_created FROM Agency a INNER JOIN Company c ON a.id_company = c.id_company")
		err = qerr
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

		return nil
	})

	return agencies, err
}
