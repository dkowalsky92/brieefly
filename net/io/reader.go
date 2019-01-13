package io

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseBody - reads from the request body into obj
func ParseBody(w http.ResponseWriter, r *http.Request, obj interface{}) {
	body, rErr := ioutil.ReadAll(r.Body)
	if rErr != nil {
		panic(rErr)
	}
	uErr := json.Unmarshal(body, &obj)
	if uErr != nil {
		panic(uErr)
	}
}
