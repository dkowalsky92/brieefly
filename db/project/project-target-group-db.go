package project

import (
	"database/sql"

	"github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/model"
)

// // TargetGroup - a target group model
// type TargetGroup struct {
// 	ID          string        `json:"id_target"`
// 	Name        string        `json:"name"`
// 	Description db.NullString `json:"description"`
// 	AgeMin      db.NullInt64  `json:"age_min"`
// 	AgeMax      db.NullInt64  `json:"age_max"`
// 	ProjectID   db.NullInt64  `json:"project"`
// }

// GetTargetGroupsForID - get project's target groups for project id
func GetTargetGroupsForID(db *db.DB, id string) ([]model.TargetGroup, *err.Error) {
	var targetGroups []model.TargetGroup

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT tg.id_target_group,
									   tg.name,
									   tg.description, 
									   tg.age_min,
									   tg.age_max
									   FROM Target_group tg 
									   INNER JOIN Project p ON p.id_project = tg.id_project
									   WHERE p.id_project = ?;`, id)
		for rows.Next() {
			var tg model.TargetGroup

			err = rows.Scan(&tg.ID,
				&tg.Name,
				&tg.Description,
				&tg.AgeMin,
				&tg.AgeMax)

			if err != nil {
				return db.HandleError(err)
			}

			targetGroups = append(targetGroups, tg)
		}

		return nil
	})

	return targetGroups, err
}
