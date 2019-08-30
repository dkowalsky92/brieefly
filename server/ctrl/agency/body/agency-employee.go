package body

import (
	"github.com/dkowalsky/brieefly/model"
)

// AgencyEmployeeBody -
type AgencyEmployeeBody struct {
	UserID    string `json:"idUser"`
	CompanyID string `json:"idCompany"`
	Role      string `json:"role"`
}

// DbAgencyEmployee -
type DbAgencyEmployee struct {
	model.AgencyEmployee
	CompanyID string `json:"idCompany" orm:"id_company"`
}

// NewDbAgencyEmployee -
// func NewDbAgencyEmployee(body AgencyEmployeeBody) *DbAgencyEmployee {
// 	return &DbAgencyEmployee{
// 		model.AgencyEmployee{
// 			UserID: body.UserID,
// 			Role:   body.Role,
// 		},
// 		body.CompanyID,
// 	}
// }

// NewDbAgencyEmployeeSingleValues -
func NewDbAgencyEmployeeSingleValues(companyid, userid string, roleid string) *DbAgencyEmployee {
	return &DbAgencyEmployee{
		model.AgencyEmployee{
			UserID: userid,
			RoleID: roleid,
		},
		companyid,
	}
}
