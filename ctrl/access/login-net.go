package access

import (
	"net/http"

	"github.com/brieefly/config"
	"github.com/brieefly/ctrl/access/body"
	"github.com/brieefly/db"
	"github.com/brieefly/net/io"
	"github.com/go-chi/chi"
)

// LoginRouter - a router with all user related routes
type LoginRouter struct {
	Mux *chi.Mux
	DB  *db.DB
}

// NewLoginRouter - creates a user subrouter to attach to the main router
func NewLoginRouter(db *db.DB) *LoginRouter {
	r := &LoginRouter{DB: db}

	mux := chi.NewRouter()
	mux.Post("/", r.Login)

	r.Mux = mux

	return r
}

// Login - logins a user and returns an auth token
func (r *LoginRouter) Login(w http.ResponseWriter, req *http.Request) {
	li := &body.LoginInfo{}
	io.ParseBody(w, req, li)
	config := config.FromContext(req.Context())
	auth, err := DbLogin(r.DB, config, li.Email, li.Password)
	io.ParseAndWrite(w, auth, err)
}
