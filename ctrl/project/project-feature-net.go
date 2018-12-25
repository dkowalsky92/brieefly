package project

import (
	"encoding/json"
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/db/project"
	"github.com/go-chi/chi"
)

type featureRouter struct {
	mux *chi.Mux
	db  *db.DB
}

func newFeatureRouter(db *db.DB) *featureRouter {
	r := &featureRouter{db: db}

	mux := chi.NewRouter()
	mux.Get("/{id}", r.getFeaturesForID)

	r.mux = mux

	return r
}

func (r *featureRouter) getFeaturesForID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	features, err := project.GetFeaturesForID(r.db, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	bytes, err := json.Marshal(features)
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
