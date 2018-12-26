package util

// SQLErr - sql related error codes
type SQLErr int64

const (
	// ErrEmptyResult - no rows returned from query
	ErrEmptyResult SQLErr = 4000
)

// Error - extended error model for better error management
type Error struct {
	Err  error
	Code int64
	Info map[string]interface{}
}

// CompositeError - a model ment to collect all the errors inbetween middleware calls
type CompositeError struct {
	errors []Error
}
