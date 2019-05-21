package body

import "github.com/brieefly/server/model"

// ProjectProcessData -
type ProjectProcessData struct {
	Phase           model.Phase                    `json:"phase"`
	Tasks           []model.Task                   `json:"tasks"`
	QuestionAnswers []ProjectQuestionAnswerOptions `json:"questions"`
}

// ProjectQuestionAnswerOptions -
type ProjectQuestionAnswerOptions struct {
	Question model.Question       `json:"question"`
	Answers  []model.AnswerOption `json:"answers"`
}
