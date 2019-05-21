package model

// CustomFeature - a custom feature model
type CustomFeature struct {
	ID          string `json:"idCustomFeature" orm:"id_custom_feature"`
	Name        string `json:"name" orm:"name"`
	Description string `json:"description" orm:"description"`
}

// -- Table: Custom_feature
// CREATE TABLE Custom_feature (
//     id_custom_feature int NOT NULL AUTO_INCREMENT,
//     name varchar(50) NOT NULL,
//     description varchar(500) NOT NULL,
//     id_project int NULL,
//     CONSTRAINT Custom_feature_pk PRIMARY KEY (id_custom_feature)
// );
