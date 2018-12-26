package model

// Feature - a custom feature model
type Feature struct {
	ID          string `json:"idFeature"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// -- Table: Feature
// CREATE TABLE Feature (
//     id_feature int NOT NULL AUTO_INCREMENT,
//     name varchar(50) NOT NULL,
//     description varchar(300) NULL,
//     CONSTRAINT Feature_pk PRIMARY KEY (id_feature)
// );
