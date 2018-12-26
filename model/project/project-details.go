package project

import (
	"github.com/brieefly/db"
	"github.com/brieefly/model"
)

// Details - project details model
type Details struct {
	ProjectID        string                 `json:"idProject"`
	Name             string                 `json:"name"`
	Type             string                 `json:"type"`
	Description      string                 `json:"description"`
	DateCreated      db.NullTime            `json:"dateCreated"`
	DateDeadline     db.NullTime            `json:"dateDeadline"`
	Status           *model.ProjectStatus   `json:"status"`
	Cms              *model.CMS             `json:"cms"`
	Colors           []model.Color          `json:"colors"`
	Features         []model.Feature        `json:"features"`
	CustomFeatures   []model.CustomFeature  `json:"customFeatures"`
	TargetGroups     []model.TargetGroup    `json:"targetGroups"`
	VisualIdentities []model.VisualIdentity `json:"visualIdentities"`
}
