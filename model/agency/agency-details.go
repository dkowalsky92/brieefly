package agency

import (
	"github.com/brieefly/model"
)

// Details - a model for agency's details and finished projects
type Details struct {
	Agency           *model.Agency   `json:"agency"`
	FinishedProjects []model.Project `json:"finishedProjects"`
}
