package model

import (
	"time"

	"github.com/brieefly/db"
)

// Project - a project model
type Project struct {
	ID               string         `json:"idProject"`
	Name             string         `json:"name"`
	Type             string         `json:"type"`
	Description      string         `json:"description"`
	OverallProgress  int64          `json:"overallProgress"`
	Language         db.NullString  `json:"language"`
	BudgetMin        db.NullInt64   `json:"budgetMin"`
	BudgetMax        db.NullInt64   `json:"budgetMax"`
	SubpageCount     db.NullInt64   `json:"subpageCount"`
	NameURL          db.NullString  `json:"nameUrl"`
	ImageURL         db.NullString  `json:"imageUrl"`
	DateDeadline     db.NullTime    `json:"dateDeadline"`
	DateCreated      time.Time      `json:"dateCreated"`
	DateLastModified db.NullTime    `json:"dateLastModified"`
	Cms              *CMS           `json:"cms"`
	Status           *ProjectStatus `json:"status"`
}
