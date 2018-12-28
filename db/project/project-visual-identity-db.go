package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// GetVisualIdentitiesForID - Get all visual identities for project id
func GetVisualIdentitiesForID(db *db.DB, id string) ([]model.VisualIdentity, error) {
	var identities []model.VisualIdentity

	err := db.WithTransaction(func(tx *sql.Tx) error {
		rows, err := tx.Query(`SELECT vi.id_visual_identity,
									   vi.type 
									   FROM Visual_identity vi 
									   WHERE vi.id_project = ?`, id)
		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		for rows.Next() {
			var vi model.VisualIdentity
			err := rows.Scan(&vi.ID,
				&vi.Type)

			if err != nil {
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			identities = append(identities, vi)
		}

		return err
	})

	return identities, err
}
