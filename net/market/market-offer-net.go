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
	mux.Get("/offers", r.getAllOffers)

	r.mux = mux

	return r
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
