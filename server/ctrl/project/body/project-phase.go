package body

import "github.com/dkowalsky/brieefly/model"

// DbPhase -
type DbPhase struct {
	model.Phase
	ProjectID string `json:"idProject" orm:"id_project"`
}

// NewDbPhase -
func NewDbPhase(phase model.Phase, projectid string) DbPhase {
	return DbPhase{
		phase,
		projectid,
	}
}
