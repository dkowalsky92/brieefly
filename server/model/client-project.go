package model

import (
	"time"

	"github.com/brieefly/server/db"
)

// ClientProject - a client-project relation model
type ClientProject struct {
	ID          string        `json:"idUser"`
	ProjectID   db.NullString `json:"idProject"`
	DateCreated time.Time     `json:"dateCreated"`
}

// -- Table: Client_project
// CREATE TABLE Client_project (
//     id_user int NOT NULL,
//     id_project int NOT NULL,
//     date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     CONSTRAINT Client_project_pk PRIMARY KEY (id_user,id_project)
// );
