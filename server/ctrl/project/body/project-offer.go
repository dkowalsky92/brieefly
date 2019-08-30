package body

import (
	"github.com/dkowalsky/brieefly/model"
)

// AgencyOffer -
type AgencyOffer struct {
	CompanyName string `json:"companyName"`
	model.Offer
}
