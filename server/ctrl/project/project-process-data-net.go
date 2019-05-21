package project

import (
	"net/http"

	"github.com/brieefly/server/db"
	"github.com/brieefly/server/net/io"
	"github.com/go-chi/chi"
)

type processDataRouter struct {
	mux *chi.Mux
	db  *db.DB
}

func newProcessDataRouter(db *db.DB) *processDataRouter {
	r := &processDataRouter{db: db}

	mux := chi.NewRouter()
	mux.Get("/{slug}", r.getProcessDataForURL)

	r.mux = mux

	return r
}

func (r *processDataRouter) getProcessDataForURL(w http.ResponseWriter, req *http.Request) {
	slug := chi.URLParam(req, "slug")
	processData, err := DbGetProcessDataForSlug(r.db, slug)
	io.ParseAndWrite(w, processData, err)
}
