package project

import (
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/net/io"
	"github.com/go-chi/chi"
)

type detailsRouter struct {
	mux *chi.Mux
	db  *db.DB
}

func newDetailsRouter(db *db.DB) *detailsRouter {
	r := &detailsRouter{db: db}

	mux := chi.NewRouter()
	mux.Get("/{slug}", r.getDetailsForURL)

	r.mux = mux

	return r
}

func (r *detailsRouter) getDetailsForURL(w http.ResponseWriter, req *http.Request) {
	slug := chi.URLParam(req, "slug")
	details, err := DbGetDetailsForURL(r.db, slug)
	io.ParseAndWrite(w, details, err)
}
