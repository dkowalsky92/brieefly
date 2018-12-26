package project

import (
	"github.com/brieefly/db"
	"github.com/brieefly/model"
)

// Details - project details model
type Details struct {
	ProjectID        string                 `json:"id_project"`
	Name             string                 `json:"name"`
	Type             string                 `json:"type"`
	Description      string                 `json:"description"`
	DateCreated      db.NullTime            `json:"date_created"`
	DateDeadline     db.NullTime            `json:"date_deadline"`
	Status           *model.ProjectStatus   `json:"status"`
	Cms              *model.CMS             `json:"cms"`
	Colors           []model.Color          `json:"colors"`
	Features         []model.Feature        `json:"features"`
	CustomFeatures   []model.CustomFeature  `json:"custom_features"`
	TargetGroups     []model.TargetGroup    `json:"target_groups"`
	VisualIdentities []model.VisualIdentity `json:"visual_identities"`
}
