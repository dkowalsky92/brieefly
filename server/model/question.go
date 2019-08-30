package model

import (
	"time"

	"github.com/dkowalsky/brieefly/db"
)

// QuestionType -
type QuestionType int

const (
	// Text - 
	Text QuestionType = 0
	// Options - 
	Options QuestionType = 1
	// MultiOptions - 
	MultiOptions QuestionType = 2
)

// Question - a model for question
type Question struct {
	ID               string      `json:"idQuestion" orm:"id_question"`
	Type             int      	 `json:"type" orm:"type"`
	Content          string      `json:"content" orm:"content"`
	Status           string      `json:"status" orm:"status"`
	DateCreated      time.Time   `json:"dateCreated" orm:"date_created"`
	DateLastModified db.NullTime `json:"dateLastModified" orm:"date_last_modified"`
}
