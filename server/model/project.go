package model

import (
	"time"

	"github.com/brieefly/server/db"
)

// Project - a project model
type Project struct {
	ID              string        `json:"idProject" orm:"id_project"`
	Name            string        `json:"name" orm:"name"`
	Type            string        `json:"type" orm:"type"`
	Description     string        `json:"description" orm:"description"`
	NameURL         string        `json:"nameUrl" orm:"url_name"`
	OverallProgress int64         `json:"overallProgress" orm:"-"`
	Language        db.NullString `json:"language" orm:"language"`
	BudgetMin       db.NullInt64  `json:"budgetMin" orm:"budget_min"`
	BudgetMax       db.NullInt64  `json:"budgetMax" orm:"budget_max"`
	SubpageCount    db.NullInt64  `json:"subpageCount" orm:"subpage_count"`

	ImageURL         db.NullString `json:"imageUrl" orm:"-"`
	DateDeadline     db.NullTime   `json:"dateDeadline" orm:"date_deadline"`
	DateCreated      time.Time     `json:"dateCreated" orm:"-"`
	DateLastModified db.NullTime   `json:"dateLastModified" orm:"-"`
	// Cms              *CMS           `json:"cms"`
	// Status           *ProjectStatus `json:"status"`
}
