package body

import (
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/model"
)

// BasicProject - a model for basic project info
type BasicProject struct {
	ID             string         `json:"idProject"`
	Name           db.NullString  `json:"name"`
	Description    db.NullString  `json:"description"`
	Type           db.NullString  `json:"type"`
	ImageURL       db.NullString  `json:"imageUrl"`
	AverageOpinion db.NullFloat64 `json:"averageOpinion"`
}

// AgencyDetails - a model for agency's details and finished projects
type AgencyDetails struct {
	Agency         *model.Agency  `json:"agency"`
	AverageOpinion db.NullFloat64 `json:"averageOpinion"`
	Projects       []BasicProject `json:"projects"`
}
