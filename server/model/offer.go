package model

import (
	"time"
)

// Offer - an offer model
type Offer struct {
	ID           string    `json:"idOffer" orm:"id_offer"`
	SalaryMin    int64     `json:"salaryMin" orm:"salary_min"`
	SalaryMax    int64     `json:"salaryMax" orm:"salary_max"`
	IsChosen     bool      `json:"isChosen" orm:"is_chosen"`
	DateStart    time.Time `json:"dateStart" orm:"date_start"`
	DateDeadline time.Time `json:"dateDeadline" orm:"date_deadline"`
	DateCreated  time.Time `json:"dateCreated" orm:"date_created"`
	ProjectID    string    `json:"idProject" orm:"id_project"`
	CompanyID    string    `json:"idCompany" orm:"id_company"`
}
