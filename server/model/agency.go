package model

import "github.com/dkowalsky/brieefly/db"

// Agency - an agency model
type Agency struct {
	NipNumber  db.NullInt64 `json:"nipNumber"`
	Company    Company      `json:"company"`
}
