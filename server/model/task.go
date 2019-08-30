package model

// Task - a model for task
type Task struct {
	ID     string `json:"idTask" orm:"id_task"`
	Name   string `json:"name" orm:"name"`
	Value  int64  `json:"value" orm:"value"`
	IsDone bool   `json:"isDone" orm:"is_done"`
}
