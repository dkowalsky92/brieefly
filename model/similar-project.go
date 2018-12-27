package model

import "github.com/brieefly/db"

// SimilarProject - a similiar project model
type SimilarProject struct {
	ID         string        `json:"idSimilarProject"`
	ProjectURL string        `json:"projectUrl"`
	ProjectID  db.NullString `json:"idProject"`
}

// -- Table: Similar_project
// CREATE TABLE Similar_project (
//     id_similar_project int NOT NULL AUTO_INCREMENT,
//     project_url varchar(300) NOT NULL,
//     id_project int NULL,
//     CONSTRAINT Similar_project_pk PRIMARY KEY (id_similar_project)
// );
