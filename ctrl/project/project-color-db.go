package project

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/model"
)

// DbGetColorsForID - Get all visual identities for project id
func DbGetColorsForID(db *db.DB, id string) ([]model.Color, *err.Error) {
	var colors []model.Color

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT c.id_color,
									  c.hex_value
									  FROM Color c 
								      WHERE c.id_project = ?`, id)
		if err != nil {
			return db.HandleError(err)
		}

		for rows.Next() {
			var c model.Color
			err := rows.Scan(&c.ID,
				&c.HexValue)
			if err != nil {
				return db.HandleError(err)
			}

			colors = append(colors, c)
		}

		return nil
	})

	return colors, err
}
