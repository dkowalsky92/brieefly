package agency

import (
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/net/io"
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

	mux.Get("/", r.GetAll)
	mux.Mount("/details", newDetailsRouter(db).mux)

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

// GetAll - get all agencies
func (r *Router) GetAll(w http.ResponseWriter, req *http.Request) {
	agencies, err := DbGetAll(r.DB)
	io.ParseAndWrite(w, agencies, err)
}
