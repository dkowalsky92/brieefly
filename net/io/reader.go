package io

import (
	"encoding/json"
	"net/http"
)

// ParseBody - reads from the request body into obj
func ParseBody(w http.ResponseWriter, r *http.Request, obj interface{}) {
	d := json.NewDecoder(r.Body)
	dErr := d.Decode(&obj)
	if dErr != nil {
		panic(dErr)
	}
}
