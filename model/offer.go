package model

import (
	"time"

	"github.com/brieefly/db"
)

// Offer - an offer model
type Offer struct {
	ID           string        `json:"id_offer"`
	Salary       int64         `json:"salary"`
	IsChosen     bool          `json:"is_chosen"`
	DateDeadline time.Time     `json:"date"`
	DateCreated  time.Time     `json:"date_created"`
	ProjectID    db.NullString `json:"id_project"`
	CompanyID    string        `json:"id_company"`
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
