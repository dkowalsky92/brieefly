package project

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/ctrl/project/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/model"
)

// DbGetProcessDataForSlug - gets the project creation process data for project slug
func DbGetProcessDataForSlug(_db *db.DB, slug string) ([]body.ProjectProcessData, *err.Error) {
	var data []body.ProjectProcessData

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {

		phases, err := DbGetPhasesForSlug(_db, slug)
		if err != nil {
			return err
		}

		for i, p := range phases {
			percentage, cErr := DbCalculatePhaseValueAsPercentage(_db, phases, i)
			if cErr != nil {
				return cErr
			}
			p.Value = int64(*percentage)

			var pcd body.ProjectProcessData
			pcd.Phase = p

			tasks, err := DbGetTasksForPhaseID(_db, p.ID)
			if err != nil {
				return err
			}
			pcd.Tasks = tasks

			questions, err := DbGetQuestionsForPhaseID(_db, p.ID)
			if err != nil {
				return err
			}
			for _, q := range questions {
				pqao := body.ProjectQuestionAnswerOptions{Question: q}
				answerOptions, err := DbGetAnswersForQuestionID(_db, q.ID)
				if err != nil {
					return err
				}

				pqao.Answers = answerOptions

				author, err := DbGetQuestionOwnerData(_db, q.ID)
				if err != nil {
					return err
				}
				pqao.Author = *author

				pcd.QuestionAnswers = append(pcd.QuestionAnswers, pqao)
			}
			data = append(data, pcd)
		}

		return nil
	})

	return data, err
}

// DbCalculatePhaseValueAsPercentage -
func DbCalculatePhaseValueAsPercentage(_db *db.DB, phases []model.Phase, selected int) (*float64, *err.Error) {
	var percentage *float64
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		selPhase := phases[selected]
		var totalValue int64
		for _, p := range phases {
			totalValue += p.Value
		}
		var tempVal float64
		tempVal = (float64(selPhase.Value) / float64(totalValue)) * 100.0
		percentage = &tempVal

		return nil
	})

	return percentage, err
}

// DbGetQuestionOwnerData -
func DbGetQuestionOwnerData(_db *db.DB, questionid string) (*body.ProjectQuestionOwnerData, *err.Error) {
	var result *body.ProjectQuestionOwnerData
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT  u.id_user,
									u.name, 
									u.surname, 
									c.image_url
									FROM Question q
									INNER JOIN Agency_employee ae ON ae.id_user = q.id_user
									INNER JOIN User u ON u.id_user = ae.id_user
									INNER JOIN Agency a ON a.id_company = ae.id_company
									INNER JOIN Company c ON c.id_company = a.id_company
									WHERE q.id_question = ?`, questionid)
		var pqod body.ProjectQuestionOwnerData
		sqlErr := row.Scan(&pqod.UserID,
			&pqod.Name,
			&pqod.Surname,
			&pqod.AgencyImageURL)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		result = &pqod

		return nil
	})

	return result, err
}

// DbGetPhasesForSlug -
func DbGetPhasesForSlug(_db *db.DB, slug string) ([]model.Phase, *err.Error) {
	var phases []model.Phase

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT pp.id_phase,
									pp.name,
									pp.is_active,
									pp.description,
									pp.value,
									pp.progress,
									pp.order_position, 
									pp.status, 
									pp.date_created FROM Phase pp
									INNER JOIN Phase p ON p.id_phase = pp.id_phase
									INNER JOIN Project pr ON pr.id_project = pp.id_project
									WHERE pr.url_name = ? 
									ORDER BY pp.order_position;`, slug)
		if err != nil {
			return _db.HandleError(err)
		}

		for rows.Next() {
			var p model.Phase

			err := rows.Scan(&p.ID,
				&p.Name,
				&p.IsActive,
				&p.Description,
				&p.Value,
				&p.Progress,
				&p.OrderPosition,
				&p.Status,
				&p.DateCreated)

			if err != nil {
				return _db.HandleError(err)
			}

			phases = append(phases, p)
		}

		return nil
	})

	return phases, err
}

// DbGetTasksForPhaseID -
func DbGetTasksForPhaseID(_db *db.DB, id string) ([]model.Task, *err.Error) {
	var tasks []model.Task

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT t.id_task,
									  t.name, 
									  t.value, 
									  t.is_done FROM Task t
									  WHERE t.id_phase = ?;`, id)
		if err != nil {
			return _db.HandleError(err)
		}

		for rows.Next() {
			var t model.Task

			err := rows.Scan(&t.ID,
				&t.Name,
				&t.Value,
				&t.IsDone)

			if err != nil {
				return _db.HandleError(err)
			}

			tasks = append(tasks, t)
		}

		return nil
	})

	return tasks, err
}

// DbGetQuestionsForPhaseID -
func DbGetQuestionsForPhaseID(_db *db.DB, id string) ([]model.Question, *err.Error) {
	var questions []model.Question

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT q.id_question, 
										q.type, 
										q.content, 
										q.status, 
										q.date_created, 
										q.date_last_modified
										FROM Question q
										WHERE q.id_phase = ?
										ORDER BY q.date_created DESC;`, id)
		if err != nil {
			return _db.HandleError(err)
		}

		for rows.Next() {
			var q model.Question

			err = rows.Scan(&q.ID,
				&q.Type,
				&q.Content,
				&q.Status,
				&q.DateCreated,
				&q.DateLastModified)

			if err != nil {
				return _db.HandleError(err)
			}

			questions = append(questions, q)
		}

		return nil
	})

	return questions, err
}

// DbInsertProcessData -
func DbInsertProcessData(_db *db.DB, bd body.ProjectProcessDataBody, projectslug string) *err.Error {
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		id, err := DbGetIDForSlug(_db, projectslug)
		if err != nil {
			return err
		}
		for _, phase := range bd.Phases {
			phaseid, pErr := DbInsertPhase(_db, phase, id.String)
			if pErr != nil {
				return pErr
			}
			for _, task := range phase.Tasks {
				tErr := DbInsertTask(_db, task, *phaseid)
				if tErr != nil {
					return tErr
				}
			}
			pErr = DbUpdatePhaseValue(_db, *phaseid)
		}

		return nil
	})

	return err
}

// DbUpdatePhaseValue -
func DbUpdatePhaseValue(_db *db.DB, phaseid string) *err.Error {
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		tasks, sqlErr := DbGetTasksForPhaseID(_db, phaseid)
		if sqlErr != nil {
			return sqlErr
		}

		var valueSum int64
		for _, task := range tasks {
			valueSum += task.Value
		}

		_, err := _db.Exec(`UPDATE Phase SET value = ? WHERE id_phase = ?`, valueSum, phaseid)
		if err != nil {
			return _db.HandleError(err)
		}

		return nil
	})

	return err
}
