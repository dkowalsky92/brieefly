package project

import (
	"encoding/json"
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/db/project"
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

	mux.Get("/{id}", r.GetAllForUserID)

	mux.Route("/name", func(sr chi.Router) {
		sr.Get("/{id}", r.GetNameForID)
	})
	mux.Mount("/cms", newCMSRouter(db).mux)
	mux.Mount("/status", newStatusRouter(db).mux)
	mux.Mount("/features", newFeatureRouter(db).mux)

	r.Mux = mux

	return r
}

// // Get - get project for id
// func (r *Router) Get(w http.ResponseWriter, req *http.Request) {
// 	id := chi.URLParam(req, "id")
// 	project, err := project.Get(r.db, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	err = json.NewEncoder(w).Encode(project)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// }

// GetAllForUserID - get all projects for user id
func (r *Router) GetAllForUserID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	projects, err := project.GetAllForUserID(r.DB, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	bytes, err := json.Marshal(projects)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// GetNameForID - get project name for project id
func (r *Router) GetNameForID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	name, err := project.GetNameForID(r.DB, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
