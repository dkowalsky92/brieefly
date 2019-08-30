package agency

import (
	"database/sql"
	"errors"

	"github.com/dkowalsky/brieefly/ctrl/agency/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/log"
)

// DbEmployeeExists -
func DbEmployeeExists(_db *db.DB, empid, companyid string) bool {
	exists := false
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT id_user, id_company FROM Agency_employee WHERE id_user = ? AND id_company = ?`, empid, companyid)
		var tempEmpID string
		var tempCompanyID string
		sqlErr := row.Scan(&tempEmpID, &tempCompanyID)
		if sqlErr != nil {
			exists = false
			return _db.HandleError(sqlErr)
		}

		exists = true
		return nil
	})

	if err != nil {
		log.Error(err)
	}

	return exists
}

// DbAddEmployee -
func DbAddEmployee(_db *db.DB, bd body.AgencyEmployeeBody) *err.Error {
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		agnExists := DbAgencyExists(_db, bd.CompanyID)
		if agnExists != true {
			return err.New(errors.New("agency does not exist, create an agency first"), 400, nil)
		}
		empExists := DbEmployeeExists(_db, bd.UserID, bd.CompanyID)
		if empExists == true {
			return err.New(errors.New("employee already exists"), 400, nil)
		}
		roleid, rErr := DbGetRoleID(_db, bd.Role)
		if rErr != nil {
			return rErr
		}
		emp := body.NewDbAgencyEmployeeSingleValues(bd.CompanyID, bd.UserID, *roleid)
		empStmt := db.InsertStmt(tx, emp, "Agency_employee")
		_, err := empStmt.Stmt.Exec(empStmt.Args...)
		if err != nil {
			return _db.HandleError(err)
		}

		return nil
	})

	return err
}
