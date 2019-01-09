package agency

import (
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/db/agency"
	"github.com/brieefly/net/io"
	"github.com/go-chi/chi"
)

type detailsRouter struct {
	mux *chi.Mux
	db  *db.DB
}

// NewRouter - creates a user subrouter to attach to the main router
func newDetailsRouter(db *db.DB) *detailsRouter {
	r := &detailsRouter{db: db}

	mux := chi.NewRouter()

	mux.Get("/{slug}", r.getDetailsForURL)
	mux.Get("/projects/{slug}", r.getFinishedProjectsForURL)

	r.mux = mux

	return r
}

func (r *detailsRouter) getFinishedProjectsForURL(w http.ResponseWriter, req *http.Request) {
	slug := chi.URLParam(req, "slug")
	finishedProjects, err := agency.GetFinishedProjectsForURL(r.db, slug)
	io.ParseAndWrite(w, finishedProjects, err)
}

func (r *detailsRouter) getDetailsForURL(w http.ResponseWriter, req *http.Request) {
	slug := chi.URLParam(req, "slug")
	details, err := agency.GetAgencyAndOpinionsForURL(r.db, slug)
	io.ParseAndWrite(w, details, err)
}
