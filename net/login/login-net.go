package login

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/brieefly/config"
	"github.com/brieefly/db"
	"github.com/brieefly/db/login"
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
	mux.Post("/", r.Login)

	r.Mux = mux

	return r
}

// Login - logins a user and returns an auth token
func (r *Router) Login(w http.ResponseWriter, req *http.Request) {
	li := login.Info{}
	body, rErr := ioutil.ReadAll(req.Body)
	if rErr != nil {
		panic(rErr)
	}
	uErr := json.Unmarshal(body, &li)
	if uErr != nil {
		panic(uErr)
	}
	config := config.FromContext(req.Context())
	auth, err := login.Login(r.DB, config, li.Email, li.Password)
	io.ParseAndWrite(w, auth, err)
}
