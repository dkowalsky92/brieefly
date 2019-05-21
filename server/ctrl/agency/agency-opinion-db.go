package agency

import (
	"database/sql"

	"github.com/brieefly/server/db"
	"github.com/brieefly/server/err"
)

// DbGetAverageOpinionForAgencyID -
func DbGetAverageOpinionForAgencyID(db *db.DB, id string) (*float64, *err.Error) {
	var averageOpinion *float64

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT AVG(oo.grade) FROM Opinion oo
									INNER JOIN Project p ON p.id_project = oo.id_project
									INNER JOIN Offer o ON o.id_project = p.id_project
									INNER JOIN Agency a ON a.id_company = o.id_company 
									WHERE a.id_company = ?;`, id)
		var ao float64
		err := row.Scan(&ao)
		if err != nil {
			return db.HandleError(err)
		}
		averageOpinion = &ao

		return nil
	})

	return averageOpinion, err
}
