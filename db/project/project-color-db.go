package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// GetColorsForID - Get all visual identities for project id
func GetColorsForID(db *db.DB, id string) ([]model.Color, error) {
	var colors []model.Color

	err := db.WithTransaction(func(tx *sql.Tx) error {
		rows, err := tx.Query(`SELECT c.id_color,
									  c.hex_value
									  FROM Color c 
								      WHERE c.id_project = ?`, id)
		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		for rows.Next() {
			var c model.Color
			err := rows.Scan(&c.ID,
				&c.HexValue)
			if err != nil {
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			colors = append(colors, c)
		}

		return err
	})

	return colors, err
}
