package body

import (
	"time"

	"github.com/dkowalsky/brieefly/model"
	"github.com/dkowalsky/brieefly/util"
)

// OfferBody -
type OfferBody struct {
	ProjectID    string      `json:"idProject"`
	BudgetMin    int64       `json:"budgetMin"`
	BudgetMax    int64       `json:"budgetMax"`
	DateDeadline time.Time `json:"dateDeadline"`
}

// NewOffer -
func NewOffer(body OfferBody, companyid string) model.Offer {
	uuid := util.UUID()
	return model.Offer{
		ID:           uuid.String(),
		SalaryMin:    body.BudgetMin,
		SalaryMax:    body.BudgetMax,
		DateCreated:  time.Now(),
		DateDeadline: body.DateDeadline,
		ProjectID:    body.ProjectID,
		CompanyID:    companyid,
		IsChosen:     false,
	}
}
