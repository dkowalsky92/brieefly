package util

import "github.com/google/uuid"

// UUID - create a new uuid for use as unique identifier
func UUID() uuid.UUID {
	return uuid.New()
}
