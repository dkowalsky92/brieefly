package net

import (
	"net/http"
	"time"

	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/ctrl/access"
	"github.com/dkowalsky/brieefly/ctrl/agency"
	"github.com/dkowalsky/brieefly/ctrl/market"
	"github.com/dkowalsky/brieefly/ctrl/project"
	"github.com/dkowalsky/brieefly/ctrl/user"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/net/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router - hub for networking
type Router struct {
	mux      *chi.Mux
	config   *config.Config
	database *db.DB
}

// NewRouter - creates a new router
func NewRouter(db *db.DB, config *config.Config) *Router {
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

		r.Mount("/api/register", access.NewRegisterRouter(db).Mux)
		r.Mount("/api/login", access.NewLoginRouter(db).Mux)
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
func (r *Router) Run() *err.Error {
	path := config.MyPath(r.config)
	httpErr := http.ListenAndServeTLS(path, r.config.TLSCert(), r.config.TLSKey(), r.mux)
	return err.New(httpErr, err.ErrInternal, nil)
}
