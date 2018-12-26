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
	mux.Get("/{id}", r.getDetailsForID)

	r.mux = mux

	return r
}

func (r *detailsRouter) getDetailsForID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	details, err := project.GetDetailsForID(r.db, id)
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
