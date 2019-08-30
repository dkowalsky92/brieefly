package project

import (
	"net/http"

	"github.com/dkowalsky/brieefly/ctrl/project/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/net/auth"
	"github.com/dkowalsky/brieefly/net/io"
	"github.com/go-chi/chi"
)

// Router - a router with all user related routes
type Router struct {
	Mux *chi.Mux
	DB  *db.DB
}

// NewRouter - creates a project subrouter to attach to the main router
func NewRouter(db *db.DB) *Router {
	r := &Router{DB: db}

	mux := chi.NewRouter()

	mux.Post("/", r.Create)
	mux.Put("/{slug}/offers/{idOffer}/choose", r.markOfferChosen)
	mux.Post("/{slug}/process", r.insertProcessData)

	mux.Route("/user", func(sr chi.Router) {
		sr.Get("/{id}", r.GetAllForUserID)
	})

	mux.Route("/name", func(sr chi.Router) {
		sr.Get("/{id}", r.GetNameForID)
	})

	mux.Mount("/cms", newCMSRouter(db).mux)
	mux.Mount("/status", newStatusRouter(db).mux)
	mux.Mount("/features", newFeatureRouter(db).mux)
	mux.Mount("/details", newDetailsRouter(db).mux)
	mux.Mount("/offers", newOfferRouter(db).mux)
	mux.Mount("/process", newProcessDataRouter(db).mux)

	mux.Mount("/questions", newQuestionAnswerRouter(db).mux)

	r.Mux = mux

	return r
}

// GetAllForUserID - get all projects for user id
func (r *Router) GetAllForUserID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	projects, err := DbGetAllForUserID(r.DB, id)
	io.ParseAndWrite(w, projects, err)
}

// GetNameForID - get project name for project id
func (r *Router) GetNameForID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	name, err := DbGetNameForID(r.DB, id)
	io.ParseAndWrite(w, name, err)
}

// Create - create project name for project id
func (r *Router) Create(w http.ResponseWriter, req *http.Request) {
	pb := &body.Body{}
	io.ParseBody(w, req, pb)
	id := auth.UserIDFromContext(req.Context())
	newProject, err := DbInsert(r.DB, *id, pb)
	io.ParseAndWrite(w, newProject, err)
}

func (r *Router) markOfferChosen(w http.ResponseWriter, req *http.Request) {
	idOffer := chi.URLParam(req, "idOffer")
	slug := chi.URLParam(req, "slug")
	err := DbMarkChosen(r.DB, idOffer, slug)
	io.WriteStatus(w, http.StatusNoContent, err)
}

func (r *Router) insertProcessData(w http.ResponseWriter, req *http.Request) {
	slug := chi.URLParam(req, "slug")
	bd := &body.ProjectProcessDataBody{}
	io.ParseBody(w, req, bd)
	err := DbInsertProcessData(r.DB, *bd, slug)
	io.WriteStatus(w, http.StatusCreated, err)
}
