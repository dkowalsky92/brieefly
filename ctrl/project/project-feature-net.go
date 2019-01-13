package project

import (
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/net/io"
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
	features, err := DbGetFeaturesForID(r.db, id)
	io.ParseAndWrite(w, features, err)
}
