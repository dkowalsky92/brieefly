package agency

import (
	"encoding/json"
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/db/agency"
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

	mux.Get("/{name}", r.getDetailsForName)

	r.mux = mux

	return r
}

// GetAll - get all users
func (r *detailsRouter) getDetailsForName(w http.ResponseWriter, req *http.Request) {
	name := chi.URLParam(req, "name")
	details, err := agency.GetDetailsForName(r.db, name)
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
