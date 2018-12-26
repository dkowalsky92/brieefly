package model

// ProjectStatus - project status model
type ProjectStatus struct {
	ID   string `json:"idStatus"`
	Name string `json:"name"`
}

// -- Table: Status
// CREATE TABLE Status (
//     id_status int NOT NULL AUTO_INCREMENT,
//     name varchar(30) NOT NULL,
//     CONSTRAINT Status_pk PRIMARY KEY (id_status)
// );
