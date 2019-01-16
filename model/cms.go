package model

import "github.com/brieefly/db"

// CMS - a cms model
type CMS struct {
	ID          string        `json:"idCms" orm:"id_cms"`
	Name        string        `json:"name" orm:"name"`
	Description db.NullString `json:"description" orm:"description"`
}

// -- Table: Cms
// CREATE TABLE Cms (
//     id_cms int NOT NULL AUTO_INCREMENT,
//     name varchar(40) NOT NULL,
//     description varchar(300) NULL,
//     CONSTRAINT Cms_pk PRIMARY KEY (id_cms)
// );
