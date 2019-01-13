package project

import (
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/net/io"
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
	status, err := DbGetStatusForID(r.db, id)
	io.ParseAndWrite(w, status, err)
}
