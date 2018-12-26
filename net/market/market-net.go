package market

import (
	"github.com/brieefly/db"
	"github.com/go-chi/chi"
)

// Router - a router with all user related routes
type Router struct {
	Mux *chi.Mux
	DB  *db.DB
}

// NewRouter - creates a project subrouter to attach to the main router
func NewRouter(db *db.DB) *Router {
	r := &Router{DB: db}

	mux := chi.NewRouter()

	mux.Mount("/", newOfferRouter(db).mux)

	r.Mux = mux

	return r
}
