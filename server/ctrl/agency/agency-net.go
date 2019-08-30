package agency

import (
	"net/http"

	"github.com/dkowalsky/brieefly/ctrl/agency/body"
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

// NewRouter - creates a user subrouter to attach to the main router
func NewRouter(db *db.DB) *Router {
	r := &Router{DB: db}

	mux := chi.NewRouter()
	mux.Post("/", r.Insert)
	mux.Get("/", r.GetAll)

	mux.Mount("/details", newDetailsRouter(db).mux)
	mux.Mount("/employees", newEmployeeRouter(db).mux)
	
	r.Mux = mux

	return r
}

// GetAll - get all agencies
func (r *Router) GetAll(w http.ResponseWriter, req *http.Request) {
	agencies, err := DbGetAll(r.DB)
	io.ParseAndWrite(w, agencies, err)
}

// Insert -
func (r *Router) Insert(w http.ResponseWriter, req *http.Request) {
	ab := &body.AgencyBody{}
	io.ParseBody(w, req, ab)
	id := auth.UserIDFromContext(req.Context())
	err := DbInsert(r.DB, *ab, *id)
	io.WriteStatus(w, http.StatusNoContent, err)
}
