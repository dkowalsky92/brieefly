package model

// Task - a model for task
type Task struct {
	ID     string `json:"idTask"`
	Name   string `json:"name"`
	Value  int64  `json:"value"`
	IsDone bool   `json:"isDone"`
}

// -- Table: Task
// CREATE TABLE Task (
//     id_task int NOT NULL AUTO_INCREMENT,
//     name varchar(30) NOT NULL,
//     value int NOT NULL DEFAULT 1,
//     is_done bool NOT NULL DEFAULT false,
//     id_project_phase int NOT NULL,
//     CONSTRAINT Task_pk PRIMARY KEY (id_task)
// );
