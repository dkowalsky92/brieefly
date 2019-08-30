package project

import (
	"net/http"

	"github.com/dkowalsky/brieefly/ctrl/project/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/net/auth"
	"github.com/dkowalsky/brieefly/net/io"
	"github.com/go-chi/chi"
)

type questionAnswerRouter struct {
	mux *chi.Mux
	db  *db.DB
}

func newQuestionAnswerRouter(db *db.DB) *questionAnswerRouter {
	r := &questionAnswerRouter{db: db}

	mux := chi.NewRouter()
	mux.Get("/types", r.getQuestionTypes)
	mux.Delete("/{question-id}", r.deleteQuestion)
	mux.Post("/", r.insertNewQuestion)
	mux.Post("/{question-id}/answer", r.insertNewAnswer)
	mux.Put("/choose", r.markAnswerChosen)

	r.mux = mux

	return r
}

func (r *questionAnswerRouter) insertNewQuestion(w http.ResponseWriter, req *http.Request) {
	bd := &body.QuestionBody{}
	io.ParseBody(w, req, bd)
	id := auth.UserIDFromContext(req.Context())
	err := DbInsertQuestion(r.db, *bd, *id)
	io.WriteStatus(w, http.StatusCreated, err)
}

func (r *questionAnswerRouter) insertNewAnswer(w http.ResponseWriter, req *http.Request) {
	questionid := chi.URLParam(req, "question-id")
	bd := &body.AnswerOptionBody{}
	io.ParseBody(w, req, bd)
	err := DbInsertAnswerOption(r.db, *bd, questionid)
	io.WriteStatus(w, http.StatusCreated, err)
}

func (r *questionAnswerRouter) deleteQuestion(w http.ResponseWriter, req *http.Request) {
	questionid := chi.URLParam(req, "question-id")
	err := DbDeleteQuestion(r.db, questionid)
	io.WriteStatus(w, http.StatusNoContent, err)
}

func (r *questionAnswerRouter) getQuestionTypes(w http.ResponseWriter, req *http.Request) {
	io.ParseAndWrite(w, body.AllQuestionTypes(), nil)
}

func (r *questionAnswerRouter) markAnswerChosen(w http.ResponseWriter, req *http.Request) {
	bd := &body.MarkAnswerChosenBody{}
	io.ParseBody(w, req, bd)
	err := DbMarkAnswerChosen(r.db, bd.QuestionID, bd.AnswerID)
	io.WriteStatus(w, http.StatusNoContent, err)
}
