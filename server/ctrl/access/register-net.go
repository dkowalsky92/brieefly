package access

import (
	"net/http"

	"github.com/dkowalsky/brieefly/ctrl/access/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/net/io"
	"github.com/go-chi/chi"
)

// RegisterRouter - a router with all user related routes
type RegisterRouter struct {
	Mux *chi.Mux
	DB  *db.DB
}

// NewRegisterRouter - creates a user subrouter to attach to the main router
func NewRegisterRouter(db *db.DB) *RegisterRouter {
	r := &RegisterRouter{DB: db}

	mux := chi.NewRouter()
	mux.Post("/", r.Register)

	r.Mux = mux

	return r
}

// Register - register a user and returns a basic user info
func (r *RegisterRouter) Register(w http.ResponseWriter, req *http.Request) {
	ru := &body.RegisterInfo{}
	io.ParseBody(w, req, ru)
	user, err := DbRegister(r.DB, ru.Email, ru.Password)
	io.ParseAndWrite(w, user, err)
}
