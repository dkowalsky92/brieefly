package model

import "github.com/brieefly/db"

// CustomFeature - a custom feature model
type CustomFeature struct {
	ID          string        `json:"id_custom_feature"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	ProjectID   db.NullString `json:"id_project"`
}

// -- Table: Custom_feature
// CREATE TABLE Custom_feature (
//     id_custom_feature int NOT NULL AUTO_INCREMENT,
//     name varchar(50) NOT NULL,
//     description varchar(500) NOT NULL,
//     id_project int NULL,
//     CONSTRAINT Custom_feature_pk PRIMARY KEY (id_custom_feature)
// );
