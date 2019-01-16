package project

import (
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/net/io"
	"github.com/go-chi/chi"
)

type offerRouter struct {
	mux *chi.Mux
	db  *db.DB
}

func newOfferRouter(db *db.DB) *offerRouter {
	r := &offerRouter{db: db}

	mux := chi.NewRouter()
	mux.Get("/{slug}", r.getOffersForSlug)

	r.mux = mux

	return r
}

func (r *offerRouter) getOffersForSlug(w http.ResponseWriter, req *http.Request) {
	slug := chi.URLParam(req, "slug")
	status, err := DbGetOffersForSlug(r.db, slug)
	io.ParseAndWrite(w, status, err)
}
