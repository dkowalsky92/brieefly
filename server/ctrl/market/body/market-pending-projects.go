package body

import (
	"github.com/brieefly/server/db"
)

// PendingProject - a model for pending projects
type PendingProject struct {
	ID                 string        `json:"idProject"`
	Name               db.NullString `json:"name"`
	Type               db.NullString `json:"type"`
	Description        db.NullString `json:"description"`
	Language           db.NullString `json:"language"`
	BudgetMin          db.NullInt64  `json:"budgetMin"`
	BudgetMax          db.NullInt64  `json:"budgetMax"`
	DateCreated        db.NullTime   `json:"dateCreated"`
	DateDeadline       db.NullTime   `json:"dateDeadline"`
	CustomFeatureCount db.NullInt64  `json:"customFeatureCount"`
	VisualIdentityType db.NullString `json:"visualIdentityType"`
}
