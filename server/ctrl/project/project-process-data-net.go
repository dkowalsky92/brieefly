package project

import (
	"net/http"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/net/io"
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
	mux.Put("/task/{task-id}/done", r.markTaskDone)

	r.mux = mux

	return r
}

func (r *processDataRouter) getProcessDataForURL(w http.ResponseWriter, req *http.Request) {
	slug := chi.URLParam(req, "slug")
	processData, err := DbGetProcessDataForSlug(r.db, slug)
	io.ParseAndWrite(w, processData, err)
}

func (r *processDataRouter) markTaskDone(w http.ResponseWriter, req *http.Request) {
	taskid := chi.URLParam(req, "task-id")
	err := DbMarkTaskDone(r.db, taskid)
	io.WriteStatus(w, http.StatusNoContent, err)
}
