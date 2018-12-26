package market

import (
	"encoding/json"
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/db/market"
	"github.com/go-chi/chi"
)

type pendingProjectsRouter struct {
	mux *chi.Mux
	db  *db.DB
}

func newPendingProjectsRouter(db *db.DB) *pendingProjectsRouter {
	r := &pendingProjectsRouter{db: db}

	mux := chi.NewRouter()
	mux.Get("/", r.getAllPendingProjects)

	r.mux = mux

	return r
}

func (r *pendingProjectsRouter) getAllPendingProjects(w http.ResponseWriter, req *http.Request) {
	projects, err := market.GetPendingProjects(r.db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	bytes, err := json.Marshal(&projects)
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
