package body

import (
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/model"
	"github.com/dkowalsky/brieefly/util"
)

// AgencyBody -
type AgencyBody struct {
	NipNumber   db.NullInt64  `json:"nipNumber" orm:"nip_number"`
	Email       string        `json:"email" orm:"email"`
	Name        db.NullString `json:"name" orm:"email"`
	Phone       db.NullString `json:"phone" orm:"phone"`
	Address     db.NullString `json:"address" orm:"address"`
	NameURL     db.NullString `json:"nameUrl" orm:"url_name"`
	WebsiteURL  db.NullString `json:"websiteUrl" orm:"website_url"`
	ImageURL    db.NullString `json:"imageUrl" orm:"image_url"`
	Description db.NullString `json:"description" orm:"description"`
}

// Company -
type Company struct {
	model.Company
}

// NewCompany -
func NewCompany(body AgencyBody) *Company {
	uuid := util.UUID().String()
	return &Company{
		model.Company{
			ID:          uuid,
			Email:       body.Email,
			Name:        body.Name,
			Phone:       body.Phone,
			Address:     body.Address,
			NameURL:     body.NameURL,
			WebsiteURL:  body.WebsiteURL,
			ImageURL:    body.ImageURL,
			Description: body.Description,
		},
	}
}

// Agency -
type Agency struct {
	CompanyID string       `json:"idCompany" orm:"id_company"`
	NipNumber db.NullInt64 `json:"nipNumber" orm:"nip_number"`
}

// NewAgency -
func NewAgency(nip db.NullInt64, companyid string) *Agency {
	return &Agency{
		NipNumber: nip,
		CompanyID: companyid,
	}
}

