package err

import (
	"encoding/json"
	"net/http"
)

// ErrorType - error
type ErrorType uint64

const (

	/*** Internal errors ***/

	// ErrConfigMalformed - config couldn't be loaded
	ErrConfigMalformed int = 3000

	/*** SQL errors ***/

	// ErrEmptyResult - no rows returned from query
	ErrEmptyResult int = 4000

	// ErrMalformedQuery - the query is malformed
	ErrMalformedQuery int = 4001

	// ErrTxBegin - transaction begin error
	ErrTxBegin int = 4002

	// ErrTxCommit - transaction commit error
	ErrTxCommit int = 4003

	// ErrConnectionFailure - transaction commit error
	ErrConnectionFailure int = 4004

	/*** Client errors ***/

	// ErrBadRequest - request sent to the server is invalid
	ErrBadRequest int = 400

	// ErrUnauthorized - the user is not authenticated
	ErrUnauthorized int = 401

	// ErrForbidden - the user is trying to access resources without sufficient permissions
	ErrForbidden int = 403

	// ErrNotFound - the destination does not exist
	ErrNotFound int = 404

	// ErrConflict - the resource posted encountered a conflict
	ErrConflict int = 409

	/*** Server errors ***/

	// ErrInternal - the server cannot process the request for an unknown reason
	ErrInternal int = 500

	// ErrBadGateway - the server is a gateway or proxy server, and it is not receiving a valid response from the backend servers that should actually fulfill the request
	ErrBadGateway int = 502

	// ErrServiceUnavailable - the server is overloaded or under maintenance. This error implies that the service should become available at some point
	ErrServiceUnavailable int = 503

	// ErrGatewayTimeout - the server is a gateway or proxy server, and it is not receiving a response from the backend servers within the allowed time period
	ErrGatewayTimeout int = 504
)

// ErrorHandler - an error handling interface with one method
type ErrorHandler interface {
	HandleError(error) *Error
}

// TypedErrorHandler - an error handling interface with specific type
type TypedErrorHandler interface {
	HandleTypedError(error, ErrorType) *Error
}

// Error - extended error model for better error management
type Error struct {
	Err  string                 `json:"error"`
	Code int                    `json:"code"`
	Info map[string]interface{} `json:"info"`
}

// WriteError - error writing to response writer
func WriteError(err *Error, w http.ResponseWriter) {
	if err == nil {
		return
	}

	switch err.Code {
	case ErrMalformedQuery:
		err.Code = ErrInternal
	case ErrEmptyResult:
		err.Code = ErrNotFound
	default:
		break
	}

	bytes, mErr := json.Marshal(err)
	if mErr != nil {
		panic(mErr)
	}

	w.WriteHeader(err.Code)

	_, wErr := w.Write(bytes)
	if wErr != nil {
		panic(wErr)
	}
}

// New - creates a new error
func New(err error, code int, info map[string]interface{}) *Error {
	return &Error{Err: err.Error(), Code: code, Info: info}
}

// CompositeError - a model ment to collect all the errors inbetween middleware calls
type CompositeError struct {
	Errors []Error `json:"errors"`
}

// Compose - creates a new CompositeError with given Error objects
func Compose(errors ...Error) *CompositeError {
	return &CompositeError{Errors: errors}
}

// Add - adds given Error to the error array
func (ce *CompositeError) Add(err Error) {
	if ce != nil {
		ce.Errors = append(ce.Errors, err)
		return
	}
	ce = Compose(err)
}

// Merge - merges two composite errors
func (ce *CompositeError) Merge(other *CompositeError) {
	if other == nil {
		return
	}

	for _, v := range other.Errors {
		ce.Errors = append(ce.Errors, v)
	}
}
