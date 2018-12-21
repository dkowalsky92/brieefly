package model

import "github.com/dkowalsky/brieefly/db"

// CMS - a cms model
type CMS struct {
	ID          int64         `json:"id_cms"`
	Name        string        `json:"name"`
	Description db.NullString `json:"description"`
}

// -- Table: Cms
// CREATE TABLE Cms (
//     id_cms int NOT NULL AUTO_INCREMENT,
//     name varchar(40) NOT NULL,
//     description varchar(300) NULL,
//     CONSTRAINT Cms_pk PRIMARY KEY (id_cms)
// );
