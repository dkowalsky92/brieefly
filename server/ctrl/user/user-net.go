package user

import (
	"net/http"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/net/auth"
	"github.com/dkowalsky/brieefly/ctrl/user/body"
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
	mux.Get("/", r.GetAll)
	mux.Put("/password", r.ChangePassword)
	mux.Put("/", r.UpdateUser)
	mux.Route("/{id}", func(sr chi.Router) {
		sr.Get("/", r.Get)
	})

	r.Mux = mux

	return r
}

// Get - get user for id
func (r *Router) Get(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	user, err := DbGet(r.DB, id)
	io.ParseAndWrite(w, user, err)
}

// GetAll - get all users
func (r *Router) GetAll(w http.ResponseWriter, req *http.Request) {
	users, err := DbGetAll(r.DB)
	io.ParseAndWrite(w, users, err)
}

// ChangePassword - change user's password
func (r *Router) ChangePassword(w http.ResponseWriter, req *http.Request) {
	pass := &body.Password{}
	io.ParseBody(w, req, pass)
	id := auth.UserIDFromContext(req.Context())
	err := DbChangePassword(r.DB, *id, pass.Password)
	io.WriteStatus(w, http.StatusNoContent, err)
}

// UpdateUser - change user's password
func (r *Router) UpdateUser(w http.ResponseWriter, req *http.Request) {
	up := &body.UserUpdate{}
	io.ParseBody(w, req, up)
	id := auth.UserIDFromContext(req.Context())
	err := DbUpdate(r.DB, *up, *id)
	io.WriteStatus(w, http.StatusNoContent, err)
}
