package model

import (
	"time"
	"github.com/dkowalsky/brieefly/db"
)

// Opinion - an opinion model
type Opinion struct {
	ID          string    `json:"idOpinion"`
	Grade       int64     `json:"grade"`
	Description db.NullString     `json:"description"`
	DateCreated time.Time `json:"dateCreated"`
}
