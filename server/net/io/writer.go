package io

import (
	"encoding/json"
	"net/http"

	_err "github.com/dkowalsky/brieefly/err"
)

// ParseAndWrite - parses an object and writes it to the ResponseWriter
func ParseAndWrite(w http.ResponseWriter, obj interface{}, err *_err.Error) {
	if err != nil {
		_err.WriteError(err, w)
		return
	}

	bytes, mErr := json.Marshal(obj)
	if mErr != nil {
		panic(mErr)
	}

	_, mErr = w.Write(bytes)
	if err != nil {
		panic(mErr)
	}
}

// WriteStatus -
func WriteStatus(w http.ResponseWriter, status int, err *_err.Error) {
	if err != nil {
		_err.WriteError(err, w)
		return
	}

	w.WriteHeader(status)
}
