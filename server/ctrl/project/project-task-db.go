package project

import (
	"database/sql"
	"errors"

	"github.com/dkowalsky/brieefly/ctrl/project/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/log"
)

// type Task struct {
// 	ID     string `json:"idTask"`
// 	Name   string `json:"name"`
// 	Value  int64  `json:"value"`
// 	IsDone bool   `json:"isDone"`
// }

// DbGetTaskForID -
func DbGetTaskForID(_db *db.DB, taskid string) (*body.DbTask, *err.Error) {
	var result *body.DbTask
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT id_task,
									name,
									value,
									is_done,
									id_phase FROM Task WHERE id_task = ?`, taskid)
		var t body.DbTask
		sqlErr := row.Scan(&t.ID,
			&t.Name,
			&t.Value,
			&t.IsDone,
			&t.PhaseID)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		result = &t

		return nil
	})

	return result, err
}

// DbTaskExists -
func DbTaskExists(_db *db.DB, taskid string) bool {
	exists := false
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT id_task FROM Task WHERE id_task = ?`, taskid)
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

// DbMarkTaskDone -
func DbMarkTaskDone(_db *db.DB, taskid string) *err.Error {
	err1 := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		exists := DbTaskExists(_db, taskid)
		if exists != true {
			return err.New(errors.New("task does not exist"), 404, nil)
		}

		stmt, sqlErr := tx.Prepare(`UPDATE Task SET is_done = true WHERE id_task = ?`)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}
		_, sqlErr = stmt.Exec(taskid)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		return nil
	})

	err2 := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		t, tErr := DbGetTaskForID(_db, taskid)
		if tErr != nil {
			return tErr
		}
		p, pErr := DbGetPhase(_db, t.PhaseID)
		if pErr != nil {
			return pErr
		}

		if (p.Status.Valid && p.Status.String == `Finished`) || p.Progress == 100 {
			switchErr := DbSwitchToNextPhase(_db, p.ID)
			if switchErr != nil {
				return switchErr
			}
		}

		return nil
	})

	if err1 != nil && err2 != nil {
		return err.New(errors.New("multiple errors occurred"), 400, map[string]interface{}{"error 1": err1, "error 2": err2})
	} else if err1 != nil {
		return err1
	}
	return err2
}

// DbInsertTask -
func DbInsertTask(_db *db.DB, bd body.TaskBody, phaseid string) *err.Error {
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		task := body.NewProcessTask(bd, phaseid)
		taskStmt := db.InsertStmt(tx, task, "Task")

		_, err := taskStmt.Stmt.Exec(taskStmt.Args...)
		if err != nil {
			return _db.HandleError(err)
		}

		return nil
	})

	return err
}
