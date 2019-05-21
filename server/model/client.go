package model

import "github.com/brieefly/server/db"

// Client - a client model
type Client struct {
	ID        string        `json:"idUser"`
	Job       db.NullString `json:"job"`
	CompanyID db.NullString `json:"idCompany"`
}

// -- Table: Client
// CREATE TABLE Client (
//     id_user int NOT NULL,
//     job varchar(40) NULL,
//     id_company int NULL,
//     CONSTRAINT Client_pk PRIMARY KEY (id_user)
// );
