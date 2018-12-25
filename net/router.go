package net

import (
	"net/http"
	"time"

	"github.com/brieefly/config"
	"github.com/brieefly/ctrl/agency"
	"github.com/brieefly/ctrl/project"
	"github.com/brieefly/ctrl/user"
	"github.com/brieefly/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router - hub for networking
type Router struct {
	mux      *chi.Mux
	config   *config.Config
	database *db.DB
}

// BrieeflyRouter - creates a new router
func BrieeflyRouter(db *db.DB, config *config.Config) *Router {
	mux := chi.NewRouter()
	mux.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Mount("/api/projects", project.NewRouter(db).Mux)
	mux.Mount("/api/users", user.NewRouter(db).Mux)
	mux.Mount("/api/agencies", agency.NewRouter(db).Mux)

	return &Router{database: db, config: config, mux: mux}
}

// Run - starts the server
func (r *Router) Run() {
	path := config.MyPath(r.config)
	err := http.ListenAndServeTLS(path, r.config.Server.Certificate, r.config.Server.Key, r.mux)
	panic(err)
}
