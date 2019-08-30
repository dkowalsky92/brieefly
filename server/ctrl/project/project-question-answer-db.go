package project

import (
	"database/sql"
	"errors"

	"github.com/dkowalsky/brieefly/ctrl/project/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/model"
)

// DbInsertQuestion -
func DbInsertQuestion(_db *db.DB, bd body.QuestionBody, userid string) *err.Error {
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		q := body.NewDbQuestion(bd, userid)
		qStmt := db.InsertStmt(tx, q, "Question")
		_, sqlErr := qStmt.Stmt.Exec(qStmt.Args...)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		return nil
	})

	return err
}

// DbInsertAnswerOption -
func DbInsertAnswerOption(_db *db.DB, bd body.AnswerOptionBody, questionid string) *err.Error {
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		q := body.NewDbAnswerOption(bd, questionid)
		qStmt := db.InsertStmt(tx, q, "Answer_option")
		_, sqlErr := qStmt.Stmt.Exec(qStmt.Args...)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		return nil
	})

	return err
}

// DbDeleteQuestion -
func DbDeleteQuestion(_db *db.DB, questionid string) *err.Error {
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		stmt, sqlErr := tx.Prepare(`DELETE FROM Question WHERE id_question = ?`)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}
		_, sqlErr = stmt.Exec(questionid)

		return _db.HandleError(sqlErr)
	})

	return err
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

// DbAnswerExists - 
func DbAnswerExists(_db *db.DB, answerid string) bool {
	exists := false
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT id_answer_option FROM Answer_option WHERE id_answer_option = ?`, answerid)
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

// DbMarkAnswerChosen - 
func DbMarkAnswerChosen(_db *db.DB, questionid, answerid string) *err.Error {
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		exists := DbAnswerExists(_db, answerid)
		if exists != true {
			return err.New(errors.New("answer does not exist"), 404, nil)
		}

		answers, qErr := DbGetAnswersForQuestionID(_db, questionid)
		if qErr != nil {
			return qErr
		}

		for _, v := range answers {
			if v.IsChosen == true {
				return err.New(errors.New("there is already a selected answer"), 400, map[string]interface{}{"answer" : v})
			}
		}

		stmt, sqlErr := tx.Prepare(`UPDATE Answer_option SET is_chosen = true WHERE id_answer_option = ?`)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}
		_, sqlErr = stmt.Exec(answerid)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}
		return nil
	})

	return err
}

