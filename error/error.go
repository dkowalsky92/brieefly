package util

// ErrorCtxKey - a key under which all errors are stored in a given *context.Context
const ErrorCtxKey string = `ErrorCtxKey`

// SQLErr - sql related error codes
type SQLErr uint16

const (
	// ErrEmptyResult - no rows returned from query
	ErrEmptyResult SQLErr = 4000
)

// Error - extended error model for better error management
type Error struct {
	Err  error                  `json:"error"`
	Code int64                  `json:"code"`
	Info map[string]interface{} `json:"info"`
}

// CompositeError - a model ment to collect all the errors inbetween middleware calls
type CompositeError struct {
	Errors []Error `json:"errors"`
}
