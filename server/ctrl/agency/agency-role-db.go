package agency

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/log"
)

// DbGetRoleID -
func DbGetRoleID(_db *db.DB, role string) (*string, *err.Error) {
	var idRole *string
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT id_agency_role FROM Agency_role WHERE role = ?`, role)

		sqlErr := row.Scan(&idRole)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		return nil
	})

	if err != nil {
		log.Error(err)
	}

	return idRole, err
}
