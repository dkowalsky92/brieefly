package model

// ProjectFeature - a custom feature model
type ProjectFeature struct {
	ID          int64    `json:"id_custom_feature"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Project     *Project `json:"project"`
}

// -- Table: Feature
// CREATE TABLE Feature (
//     id_feature int NOT NULL AUTO_INCREMENT,
//     name varchar(50) NOT NULL,
//     description varchar(300) NULL,
//     CONSTRAINT Feature_pk PRIMARY KEY (id_feature)
// );
