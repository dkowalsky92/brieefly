package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/server/ctrl/project/body"
	"github.com/brieefly/server/db"
	"github.com/brieefly/server/err"
	"github.com/brieefly/server/model"
)

// DbGetProcessDataForSlug - gets the project creation process data for project slug
func DbGetProcessDataForSlug(_db *db.DB, id string) ([]body.ProjectProcessData, *err.Error) {
	var data []body.ProjectProcessData

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {

		phases, err := DbGetPhasesForID(_db, id)
		if err != nil {
			return err
		}

		for _, p := range phases {

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
			fmt.Printf("%+v\n\n", questions)
			for _, q := range questions {
				fmt.Printf("%+v\n\n", q)
				pqao := body.ProjectQuestionAnswerOptions{Question: q}
				answerOptions, err := DbGetAnswersForQuestionID(_db, q.ID)
				if err != nil {
					return err
				}
				fmt.Printf("%+v\n", answerOptions)
				pqao.Answers = answerOptions
				pcd.QuestionAnswers = append(pcd.QuestionAnswers, pqao)
			}
			data = append(data, pcd)
		}

		return nil
	})

	return data, err
}

// DbGetPhasesForID -
func DbGetPhasesForID(_db *db.DB, id string) ([]model.Phase, *err.Error) {
	var phases []model.Phase

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT pp.id_phase,
									pp.name,
									pp.description,
									pp.value,
									pp.progress,
									pp.order_position, 
									pp.status, 
									pp.date_created FROM Phase pp
									INNER JOIN Phase p ON p.id_phase = pp.id_phase
									INNER JOIN Project pr ON pr.id_project = pp.id_project
									WHERE pr.name_slug = ?;`, id)
		if err != nil {
			return _db.HandleError(err)
		}

		for rows.Next() {
			var p model.Phase

			err := rows.Scan(&p.ID,
				&p.Name,
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
		fmt.Printf("%+s\n\n", id)
		rows, err := tx.Query(`SELECT t.id_task,
									  t.name, 
									  t.value, 
									  t.is_done FROM Task t
									  WHERE t.id_phase = ?;`, id)
		if err != nil {
			fmt.Printf("%+v", err)
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
										q.date_last_modified, 
										q.id_user FROM Question q
										WHERE q.id_phase = ?;`, id)
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
				&q.DateLastModified,
				&q.UserID)

			if err != nil {
				return _db.HandleError(err)
			}

			fmt.Printf("%+v\n", q)
			questions = append(questions, q)
		}

		return nil
	})

	return questions, err
}

// DbGetAnswersForQuestionID -
func DbGetAnswersForQuestionID(_db *db.DB, id string) ([]model.AnswerOption, *err.Error) {
	var answers []model.AnswerOption

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		rows, err := tx.Query(`SELECT a.id_answer_option,
										a.content, 
										a.is_chosen FROM Answer_option a
										WHERE a.id_question = ?;`, id)
		if err != nil {
			return _db.HandleError(err)
		}

		for rows.Next() {
			var a model.AnswerOption

			err = rows.Scan(&a.ID,
				&a.Content,
				&a.IsChosen)

			if err != nil {
				return _db.HandleError(err)
			}

			answers = append(answers, a)
		}

		return nil
	})

	return answers, err
}
