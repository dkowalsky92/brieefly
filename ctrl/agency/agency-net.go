package agency

import (
	"encoding/json"
	"net/http"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/db/agency"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Get("/", r.GetAll)

	// mainMux.Route("/{id}", func(sr chi.Router) {
	// 	sr.Get("/", r.Get)
	// })

	r.Mux = mux

	return r
}

// // Get - get user for id
// func (r *Router) Get(w http.ResponseWriter, req *http.Request) {
// 	id := chi.URLParam(req, "id")
// 	agency, err := agency.Get(r.db, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	err = json.NewEncoder(w).Encode(user)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// }

// GetAll - get all users
func (r *Router) GetAll(w http.ResponseWriter, req *http.Request) {
	agencies, err := agency.GetAll(r.DB)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	bytes, err := json.Marshal(agencies)
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
