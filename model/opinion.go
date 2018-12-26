package model

import (
	"time"

	"github.com/brieefly/db"
)

// Opinion - an opinion model
type Opinion struct {
	ID          string        `json:"id_opinion"`
	Grade       int64         `json:"grade"`
	Description int64         `json:"description"`
	DateCreated time.Time     `json:"date_created"`
	ProjectID   db.NullString `json:"id_project"`
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
