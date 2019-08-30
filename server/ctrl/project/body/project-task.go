package body

import (
	"github.com/dkowalsky/brieefly/model"
)

// DbTask -
type DbTask struct {
	model.Task
	PhaseID string `json:"idPhase" orm:"id_phase"`
}

// NewDbTask -
func NewDbTask(task model.Task, phaseid string) DbTask {
	return DbTask{
		task,
		phaseid,
	}
}
