package model

// AnswerOption -
type AnswerOption struct {
	ID       string `json:"idAnswerOption"`
	Content  string `json:"content"`
	IsChosen bool   `json:"isChosen"`
}

// -- Table: Client_project
// CREATE TABLE Client_project (
//     id_user int NOT NULL,
//     id_project int NOT NULL,
//     date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     CONSTRAINT Client_project_pk PRIMARY KEY (id_user,id_project)
// );
