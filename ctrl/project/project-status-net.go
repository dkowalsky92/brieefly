package project

import (
	"encoding/json"
	"net/http"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/db/project"
	"github.com/go-chi/chi"
)

type statusRouter struct {
	mux *chi.Mux
	db  *db.DB
}

func newStatusRouter(db *db.DB) *statusRouter {
	r := &statusRouter{db: db}

	mux := chi.NewRouter()
	mux.Get("/{id}", r.getStatusForID)

	r.mux = mux

	return r
}

// GetStatusForID - get project status for project id
func (r *statusRouter) getStatusForID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	status, err := project.GetStatusForID(r.db, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	bytes, err := json.Marshal(status)
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
