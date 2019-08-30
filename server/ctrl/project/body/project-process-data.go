package body

import (
	"time"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/model"
	"github.com/dkowalsky/brieefly/util"
)

// ProjectProcessData -
type ProjectProcessData struct {
	Phase           model.Phase                    `json:"phase"`
	Tasks           []model.Task                   `json:"tasks"`
	QuestionAnswers []ProjectQuestionAnswerOptions `json:"questions"`
}

// ProjectQuestionAnswerOptions -
type ProjectQuestionAnswerOptions struct {
	Author   ProjectQuestionOwnerData `json:"author"`
	Question model.Question           `json:"question"`
	Answers  []model.AnswerOption     `json:"answers"`
}

// ProjectQuestionOwnerData -
type ProjectQuestionOwnerData struct {
	UserID         string `json:"idUser"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	AgencyImageURL string `json:"imageUrl"`
}

// ProjectProcessDataBody -
type ProjectProcessDataBody struct {
	Phases []PhaseBody `json:"phases"`
}

// PhaseBody -
type PhaseBody struct {
	Name          string        `json:"name"`
	Description   db.NullString `json:"description"`
	OrderPosition int64         `json:"orderPosition"`
	Tasks         []TaskBody    `json:"tasks"`
}

// TaskBody -
type TaskBody struct {
	ID    string `json:"idTask"`
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

// NewProcessPhase -
func NewProcessPhase(body PhaseBody, projectid string) DbPhase {
	uuid := util.UUID().String()
	var status string
	var isActive bool
	if body.OrderPosition == 0 {
		status = "In Progress"
		isActive = true
	} else {
		status = "Pending"
		isActive = false
	}

	return DbPhase{
		model.Phase{
			ID:            uuid,
			Name:          body.Name,
			IsActive:      isActive,
			Description:   body.Description,
			Value:         0,
			Progress:      0,
			OrderPosition: db.NullInt64{Int64: body.OrderPosition, Valid: true},
			Status:        db.NullString{String: status, Valid: true},
			DateCreated:   time.Now(),
		},
		projectid,
	}
}

// NewProcessTask -
func NewProcessTask(body TaskBody, phaseid string) DbTask {
	return DbTask{
		model.Task{
			ID:     body.ID,
			Name:   body.Name,
			Value:  body.Value,
			IsDone: false,
		},
		phaseid,
	}
}

// // Task - a model for task
// type Task struct {
// 	ID     string `json:"idTask"`
// 	Name   string `json:"name"`
// 	Value  int64  `json:"value"`
// 	IsDone bool   `json:"isDone"`
// }

// type Phase struct {
// 	ID            string        `json:"idProjectPhase"`
// 	Name          string        `json:"name"`
// 	IsActive 	  bool			`json:"isActive"`
// 	Description   db.NullString `json:"description"`
// 	Value         int64         `json:"value"`
// 	Progress      int64         `json:"progress"`
// 	OrderPosition db.NullInt64  `json:"orderPosition"`
// 	Status        db.NullString `json:"status"`
// 	DateCreated   time.Time     `json:"dateCreated"`
// }
