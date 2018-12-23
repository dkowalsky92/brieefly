package model

import "github.com/dkowalsky/brieefly/db"

// VisualIdentity - visual identity model
type VisualIdentity struct {
	ID        string        `json:"id_visual_identity"`
	Type      string        `json:"type"`
	ProjectID db.NullString `json:"id_project"`
}

// -- Table: Visual_identity
// CREATE TABLE Visual_identity (
//     id_visual_identity int NOT NULL AUTO_INCREMENT,
//     type varchar(30) NOT NULL,
//     id_project int NULL,
//     CONSTRAINT Visual_identity_pk PRIMARY KEY (id_visual_identity)
// );
