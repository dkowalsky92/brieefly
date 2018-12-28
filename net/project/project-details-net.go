package project

import (
	"encoding/json"
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/db/project"
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
	details, err := project.GetDetailsForURL(r.db, slug)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	bytes, err := json.Marshal(&details)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
