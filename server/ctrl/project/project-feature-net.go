package project

import (
	"net/http"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/net/io"
	"github.com/go-chi/chi"
)

type featureRouter struct {
	mux *chi.Mux
	db  *db.DB
}

func newFeatureRouter(db *db.DB) *featureRouter {
	r := &featureRouter{db: db}

	mux := chi.NewRouter()
	mux.Get("/", r.getAllFeatures)
	mux.Get("/{id}", r.getFeaturesForID)

	r.mux = mux

	return r
}

func (r *featureRouter) getAllFeatures(w http.ResponseWriter, req *http.Request) {
	features, err := DbGetAllFeatures(r.db)
	io.ParseAndWrite(w, features, err)
}

func (r *featureRouter) getFeaturesForID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	features, err := DbGetFeaturesForID(r.db, id)
	io.ParseAndWrite(w, features, err)
}
