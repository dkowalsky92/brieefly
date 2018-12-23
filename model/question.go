package model

import (
	"time"

	"github.com/dkowalsky/brieefly/db"
)

// Question - a model for question
type Question struct {
	ID               string      `json:"id_question"`
	Type             string      `json:"type"`
	Content          string      `json:"content"`
	Status           string      `json:"status"`
	DateCreated      time.Time   `json:"date_created"`
	DateLastModified db.NullTime `json:"date_last_modified"`
	UserID           string      `json:"id_user"`
	ProjectPhaseID   string      `json:"id_project_phase"`
}

// -- Table: Question
// CREATE TABLE Question (
//     id_question int NOT NULL AUTO_INCREMENT,
//     type varchar(100) NOT NULL,
//     content varchar(500) NOT NULL,
//     status varchar(30) NOT NULL DEFAULT 'pending',
//     date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     date_last_modified timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
//     id_user int NOT NULL,
//     id_project_phase int NOT NULL,
//     CONSTRAINT Question_pk PRIMARY KEY (id_question)
// );
