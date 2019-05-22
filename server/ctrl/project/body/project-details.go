package body

import (
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/model"
)

// ProjectDetails - project details model
type ProjectDetails struct {
	ProjectID        string                 `json:"idProject"`
	Name             string                 `json:"name"`
	Type             string                 `json:"type"`
	Description      string                 `json:"description"`
	DateCreated      db.NullTime            `json:"dateCreated"`
	DateDeadline     db.NullTime            `json:"dateDeadline"`
	AverageOpinion   float64                `json:"averageOpinion"`
	AgencyName       db.NullString          `json:"agencyName"`
	AgencyLogoURL    db.NullString          `json:"agencyImageUrl"`
	Status           *model.ProjectStatus   `json:"status"`
	Cms              *model.CMS             `json:"cms"`
	Colors           []model.Color          `json:"colors"`
	Features         []model.Feature        `json:"features"`
	CustomFeatures   []model.CustomFeature  `json:"customFeatures"`
	TargetGroups     []model.TargetGroup    `json:"targetGroups"`
	VisualIdentities []model.VisualIdentity `json:"visualIdentities"`
	SimilarProjects  []model.SimilarProject `json:"similarProjects"`
}
