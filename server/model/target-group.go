package model

import "github.com/dkowalsky/brieefly/db"

// TargetGroup - a target group model
type TargetGroup struct {
	ID          string        `json:"idTarget" orm:"id_target_group"`
	Name        string        `json:"name" orm:"name"`
	Description db.NullString `json:"description" orm:"description"`
	AgeMin      db.NullInt64  `json:"ageMin" orm:"age_min"`
	AgeMax      db.NullInt64  `json:"ageMax" orm:"age_max"`
}

// -- Table: Target_group
// CREATE TABLE Target_group (
//     id_target int NOT NULL AUTO_INCREMENT,
//     name varchar(50) NOT NULL,
//     description varchar(500) NULL,
//     age_min int NULL,
//     age_max int NULL,
//     id_project int NULL,
//     CONSTRAINT Target_group_pk PRIMARY KEY (id_target)
// );
