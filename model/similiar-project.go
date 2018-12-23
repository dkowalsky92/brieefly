package model

import "github.com/dkowalsky/brieefly/db"

// SimiliarProject - a similiar project model
type SimiliarProject struct {
	ID         string        `json:"id_similiar_project"`
	ProjectURL string        `json:"project_url"`
	ProjectID  db.NullString `json:"id_project"`
}

// -- Table: Similar_project
// CREATE TABLE Similar_project (
//     id_similar_project int NOT NULL AUTO_INCREMENT,
//     project_url varchar(300) NOT NULL,
//     id_project int NULL,
//     CONSTRAINT Similar_project_pk PRIMARY KEY (id_similar_project)
// );
