package body

import (
	"time"

	"github.com/dkowalsky/brieefly/util"
	"github.com/dkowalsky/brieefly/model"
	"github.com/dkowalsky/brieefly/db"
)
// QuestionBody - 
type QuestionBody struct {
	Content string `json:"content"`
	Type    int    `json:"type"`
	PhaseID string `json:"idPhase"`
}

// DbQuestion - 
type DbQuestion struct {
	UserID  string `json:"idUser" orm:"id_user"`
	PhaseID string `json:"idPhase" orm:"id_phase"`
	model.Question
}

// NewDbQuestion - 
func NewDbQuestion(body QuestionBody, userid string) (*DbQuestion) {
	time := time.Now()
	uuid := util.UUID()
	return &DbQuestion {
		userid,
		body.PhaseID,
		model.Question {
			ID: uuid.String(),
			Content: body.Content,
			Type: body.Type,
			Status: `Pending`,
			DateCreated: time,
			DateLastModified: db.NullTime{Time: time, Valid: true},
		},
	}
}

// AnswerOptionBody - 
type AnswerOptionBody struct {
	Content string `json:"content" orm:"content"`
	IsChosen bool `json:"isChosen" orm:"is_chosen"`
	ImageURL db.NullString `json:"imageUrl" orm:"image_url"`
}

// DbAnswerOption - 
type DbAnswerOption struct {
	QuestionID string `json:"idQuestion" orm:"id_question"`
	model.AnswerOption
}

// NewDbAnswerOption - 
func NewDbAnswerOption(body AnswerOptionBody, questionid string) (*DbAnswerOption) {
	uuid := util.UUID()
	return &DbAnswerOption {
		questionid,
		model.AnswerOption {
			ID: uuid.String(),
			Content: body.Content,
			IsChosen: body.IsChosen,
			ImageURL: body.ImageURL,
		},
	}
}

// MarkAnswerChosenBody - 
type MarkAnswerChosenBody struct {
	AnswerID string `json:"idAnswer"`
	QuestionID string `json:"idQuestion"`
}

// QuestionType - 
type QuestionType struct {
	Value model.QuestionType `json:"value"`
	Label string 			`json:"label"`
}

// NewQuestionType -
func NewQuestionType(t model.QuestionType) QuestionType {
	switch t {
	case model.Text:
		return QuestionType{Value: t, Label:"Text"}
	case model.Options:
		return QuestionType{Value: t, Label:"Options"}
	case model.MultiOptions:
		return QuestionType{Value: t, Label:"Multiple Options"}
	default:
		return QuestionType{Value: t, Label:"Text"}
	}
}

// AllQuestionTypes - 
func AllQuestionTypes() []QuestionType { 
	return []QuestionType{NewQuestionType(model.Text), NewQuestionType(model.Options), NewQuestionType(model.MultiOptions)}
}