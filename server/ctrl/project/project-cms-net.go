package project

import (
	"net/http"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/net/io"
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
	mux.Get("/", r.getAllCMS)
	
	r.mux = mux

	return r
}

// getCMSForID - get project's cms for project id
func (r *cmsRouter) getCMSForID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	cms, err := DbGetCMSForID(r.db, id)
	io.ParseAndWrite(w, cms, err)
}

// getAllCMS -
func (r *cmsRouter) getAllCMS(w http.ResponseWriter, req *http.Request) {
	cms, err := DbGetAllCMS(r.db)
	io.ParseAndWrite(w, cms, err)
}