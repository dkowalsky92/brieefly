package user

import (
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/db/user"
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
	mux.Route("/{id}", func(sr chi.Router) {
		sr.Get("/", r.Get)
	})
	r.Mux = mux

	return r
}

// Get - get user for id
func (r *Router) Get(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	user, err := user.Get(r.DB, id)
	io.ParseAndWrite(w, user, err)
}

// GetAll - get all users
func (r *Router) GetAll(w http.ResponseWriter, req *http.Request) {
	users, err := user.GetAll(r.DB)
	io.ParseAndWrite(w, users, err)
}
