package model

import "github.com/dkowalsky/brieefly/db"

// AnswerOption -
type AnswerOption struct {
	ID       string `json:"idAnswerOption" orm:"id_answer_option"`
	Content  string `json:"content" orm:"content"`
	IsChosen bool   `json:"isChosen" orm:"is_chosen"`
	ImageURL db.NullString  `json:"imageUrl" orm:"image_url"`
}
