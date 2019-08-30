package project

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/ctrl/project/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/log"
)

// DbPhaseExists -
func DbPhaseExists(_db *db.DB, phaseid string) bool {
	exists := false
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT id_phase FROM Phase WHERE id_phase = ?`, phaseid)
		var id string
		sqlErr := row.Scan(&id)
		if sqlErr != nil {
			if sqlErr == sql.ErrNoRows {
				exists = false
				return _db.HandleError(sqlErr)
			}
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

// DbGetPhase -
func DbGetPhase(_db *db.DB, phaseid string) (*body.DbPhase, *err.Error) {
	var result *body.DbPhase
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT id_phase, 
									name, 
									is_active, 
									description,
									value,
									progress,
									order_position,
									status,
									date_created,
									id_project
									FROM Phase WHERE id_phase = ?`, phaseid)
		var p body.DbPhase
		sqlErr := row.Scan(&p.ID,
			&p.Name,
			&p.IsActive,
			&p.Description,
			&p.Value,
			&p.Progress,
			&p.OrderPosition,
			&p.Status,
			&p.DateCreated,
			&p.ProjectID)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		result = &p

		return nil
	})

	return result, err
}

// DbGetNextPhase -
func DbGetNextPhase(_db *db.DB, phase body.DbPhase) (*body.DbPhase, *err.Error) {
	var result *body.DbPhase
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT id_phase, 
									name, 
									is_active, 
									description,
									value,
									progress,
									order_position,
									status,
									date_created,
									id_project
									FROM Phase WHERE order_position = ? AND id_project = ?`, phase.OrderPosition.Int64+1, phase.ProjectID)
		var p body.DbPhase
		sqlErr := row.Scan(&p.ID,
			&p.Name,
			&p.IsActive,
			&p.Description,
			&p.Value,
			&p.Progress,
			&p.OrderPosition,
			&p.Status,
			&p.DateCreated,
			&p.ProjectID)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		result = &p

		return nil
	})

	return result, err
}

// DbSwitchToNextPhase -
func DbSwitchToNextPhase(_db *db.DB, phaseid string) *err.Error {
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		currP, currPErr := DbGetPhase(_db, phaseid)
		if currPErr != nil {
			return currPErr
		}

		nextP, nextPErr := DbGetNextPhase(_db, *currP)
		if nextPErr != nil {
			return nextPErr
		}

		stmt, sqlErr := tx.Prepare(`UPDATE Phase SET is_active = true, status = ? WHERE id_phase = ?`)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		_, sqlErr = stmt.Exec(`In Progress`, nextP.ID)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		return nil
	})

	return err
}

// DbInsertPhase -
func DbInsertPhase(_db *db.DB, bd body.PhaseBody, projectid string) (*string, *err.Error) {
	var id *string
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		phase := body.NewProcessPhase(bd, projectid)
		phaseStmt := db.InsertStmt(tx, phase, "Phase")

		_, err := phaseStmt.Stmt.Exec(phaseStmt.Args...)
		if err != nil {
			return _db.HandleError(err)
		}

		id = &phase.ID

		return nil
	})

	return id, err
}
