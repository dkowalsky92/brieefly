package model

import (
	"github.com/brieefly/db"
)

// Offer - an offer model
type Offer struct {
	ID           string        `json:"idOffer"`
	Salary       int64         `json:"salary"`
	IsChosen     bool          `json:"isChosen"`
	DateDeadline db.NullTime   `json:"dateDeadline"`
	DateCreated  db.NullTime   `json:"dateCreated"`
	ProjectID    db.NullString `json:"idProject"`
	CompanyID    string        `json:"idCompany"`
}
