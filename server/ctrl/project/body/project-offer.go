package body

import "github.com/brieefly/server/model"

// AgencyOffer -
type AgencyOffer struct {
	CompanyName string `json:"companyName"`
	model.Offer
}
