package project

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/model"
)

// GetVisualIdentitiesForID - Get all visual identities for project id
func GetVisualIdentitiesForID(db *db.DB, id string) ([]model.VisualIdentity, *err.Error) {
	var identities []model.VisualIdentity

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT vi.id_visual_identity,
									   vi.type 
									   FROM Visual_identity vi 
									   WHERE vi.id_project = ?`, id)
		if err != nil {
			return db.HandleError(err)
		}

		for rows.Next() {
			var vi model.VisualIdentity
			err := rows.Scan(&vi.ID,
				&vi.Type)

			if err != nil {
				return db.HandleError(err)
			}

			identities = append(identities, vi)
		}

		return nil
	})

	return identities, err
}
