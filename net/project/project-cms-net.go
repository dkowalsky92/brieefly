package project

import (
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/db/project"
	"github.com/brieefly/net/io"
	"github.com/go-chi/chi"
)

// CMSRouter - a router with cms related routes
type cmsRouter struct {
	mux *chi.Mux
	db  *db.DB
}

// NewCMSRouter - creates a user subrouter to attach to the main router
func newCMSRouter(db *db.DB) *cmsRouter {
	r := &cmsRouter{db: db}

	mux := chi.NewRouter()
	mux.Get("/{id}", r.getCMSForID)

	r.mux = mux

	return r
}

// GetCMSForID - get project's cms for project id
func (r *cmsRouter) getCMSForID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	cms, err := project.GetCMSForID(r.db, id)
	io.ParseAndWrite(w, cms, err)
}
