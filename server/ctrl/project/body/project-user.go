package body

import (
	"github.com/brieefly/server/db"
	"github.com/brieefly/server/model"
)

// UserProject - a model for user's project
type UserProject struct {
	Project    model.Project        `json:"project"`
	CMS        *model.CMS           `json:"cms"`
	Status     *model.ProjectStatus `json:"status"`
	AgencyName db.NullString        `json:"agencyName"`
}
