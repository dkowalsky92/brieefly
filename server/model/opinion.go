package model

import (
	"time"
)

// Opinion - an opinion model
type Opinion struct {
	ID          string    `json:"idOpinion"`
	Grade       int64     `json:"grade"`
	Description int64     `json:"description"`
	DateCreated time.Time `json:"dateCreated"`
}

// -- Table: Opinion
// CREATE TABLE Opinion (
//     id_opinion varchar(100) NOT NULL,
//     grade int NOT NULL,
//     description varchar(500) NULL,
//     date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     id_project varchar(100) NOT NULL,
//     CONSTRAINT Opinion_pk PRIMARY KEY (id_opinion)
// );
