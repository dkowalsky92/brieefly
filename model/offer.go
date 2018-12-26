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

// -- Table: Offer
// CREATE TABLE Offer (
//     id_offer int NOT NULL AUTO_INCREMENT,
//     salary int NOT NULL,
//     is_chosen bool NOT NULL DEFAULT false,
//     date_deadline date NOT NULL,
//     date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     id_project int NULL,
//     id_company int NOT NULL,
//     CONSTRAINT Offer_pk PRIMARY KEY (id_offer)
// );
