package agency

import (
	"net/http"

	"github.com/dkowalsky/brieefly/ctrl/agency/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/net/io"
	"github.com/go-chi/chi"
)

type employeeRouter struct {
	mux *chi.Mux
	db  *db.DB
}

// NewRouter - creates a user subrouter to attach to the main router
func newEmployeeRouter(db *db.DB) *employeeRouter {
	r := &employeeRouter{db: db}

	mux := chi.NewRouter()
	mux.Post("/", r.addEmployee)

	r.mux = mux

	return r
}

func (r *employeeRouter) addEmployee(w http.ResponseWriter, req *http.Request) {
	bd := &body.AgencyEmployeeBody{}
	io.ParseBody(w, req, bd)
	err := DbAddEmployee(r.db, *bd)
	io.WriteStatus(w, http.StatusNoContent, err)
}
