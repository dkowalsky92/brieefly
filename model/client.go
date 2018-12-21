package model

import "github.com/dkowalsky/brieefly/db"

// Client - a client model
type Client struct {
	ID        string        `json:"id_user"`
	Job       db.NullString `json:"job"`
	CompanyID db.NullString `json:"id_company"`
}

// -- Table: Client
// CREATE TABLE Client (
//     id_user int NOT NULL,
//     job varchar(40) NULL,
//     id_company int NULL,
//     CONSTRAINT Client_pk PRIMARY KEY (id_user)
// );
