package net

import (
	"net/http"
	"time"

	"github.com/brieefly/config"
	"github.com/brieefly/db"
	"github.com/brieefly/net/agency"
	"github.com/brieefly/net/auth"
	"github.com/brieefly/net/login"
	"github.com/brieefly/net/market"
	"github.com/brieefly/net/project"
	"github.com/brieefly/net/user"
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

	/** public **/
	public(mux, db, config)

	/** protected **/
	protected(mux, db, config)

	/** private **/
	private(mux, db, config)

	return &Router{database: db, config: config, mux: mux}
}

func public(mux *chi.Mux, db *db.DB, config *config.Config) {
	mux.Group(func(r chi.Router) {
		injectMiddlewareStack(r, config)

		r.Mount("/api/login", login.NewRouter(db).Mux)
	})
}

func protected(mux *chi.Mux, db *db.DB, config *config.Config) {
	mux.Group(func(r chi.Router) {
		injectMiddlewareStack(r, config)

		r.Use(auth.ValidateTokenMiddleware)

		r.Mount("/api/projects", project.NewRouter(db).Mux)
		r.Mount("/api/users", user.NewRouter(db).Mux)
		r.Mount("/api/agencies", agency.NewRouter(db).Mux)
		r.Mount("/api/market", market.NewRouter(db).Mux)
	})
}

func private(mux *chi.Mux, db *db.DB, config *config.Config) {
	mux.Group(func(r chi.Router) {
		injectMiddlewareStack(r, config)

		r.Use(auth.ValidateTokenMiddleware)
	})
}

func injectMiddlewareStack(r chi.Router, cnf *config.Config) {
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := config.IntoContext(r.Context(), cnf)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

// Run - starts the server
func (r *Router) Run() {
	path := config.MyPath(r.config)
	err := http.ListenAndServeTLS(path, r.config.Server.Certificate, r.config.Server.Key, r.mux)
	panic(err)
}
