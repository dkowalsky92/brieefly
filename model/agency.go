package model

import "github.com/brieefly/db"

// Agency - an agency model
type Agency struct {
	AgencyCode string       `json:"agencyCode"`
	NipNumber  db.NullInt64 `json:"nipNumber"`
	Company    Company      `json:"company"`
}

// CREATE TABLE Agency (
//     id_company int NOT NULL,
//     agency_code varchar(13) NOT NULL,
//     nip_number int NULL,
//     CONSTRAINT Agency_pk PRIMARY KEY (id_company)
// );
