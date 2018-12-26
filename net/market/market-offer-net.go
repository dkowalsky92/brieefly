package market

import (
	"encoding/json"
	"net/http"

	"github.com/brieefly/db"
	"github.com/brieefly/db/market"
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
	offers, err := market.GetOffersForID(r.db, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	bytes, err := json.Marshal(&offers)
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

func (r *offerRouter) getAllOffers(w http.ResponseWriter, req *http.Request) {
	offers, err := market.GetAllOffers(r.db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	bytes, err := json.Marshal(&offers)
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
