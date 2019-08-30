package model

// AgencyEmployee -
type AgencyEmployee struct {
	UserID string `json:"id_user" orm:"id_user"`
	RoleID string `json:"id_agency_role" orm:"id_agency_role"`
}
