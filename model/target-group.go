package model

import "github.com/brieefly/db"

// TargetGroup - a target group model
type TargetGroup struct {
	ID          string        `json:"idTarget"`
	Name        string        `json:"name"`
	Description db.NullString `json:"description"`
	AgeMin      db.NullInt64  `json:"ageMin"`
	AgeMax      db.NullInt64  `json:"ageMax"`
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
