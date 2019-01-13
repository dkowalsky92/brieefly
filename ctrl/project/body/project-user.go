package body

import (
	"github.com/brieefly/db"
	"github.com/brieefly/model"
)

// UserProject - a model for user's project
type UserProject struct {
	Project    model.Project `json:"project"`
	AgencyName db.NullString `json:"agencyName"`
}
