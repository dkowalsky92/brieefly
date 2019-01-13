package market

import (
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/net/io"
	"github.com/go-chi/chi"
)

type offerRouter struct {
	mux *chi.Mux
	db  *db.DB
}

func newOfferRouter(db *db.DB) *offerRouter {
	r := &offerRouter{db: db}

	mux := chi.NewRouter()
	mux.Get("/{id}", r.getOffersForID)

	r.mux = mux

	return r
}

func (r *offerRouter) getOffersForID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	offers, err := DbGetOffersForID(r.db, id)
	io.ParseAndWrite(w, offers, err)
}

func (r *offerRouter) getAllOffers(w http.ResponseWriter, req *http.Request) {
	offers, err := DbGetAllOffers(r.db)
	io.ParseAndWrite(w, offers, err)
}
