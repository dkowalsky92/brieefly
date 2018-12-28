package agency

import (
	"github.com/brieefly/db"
	"github.com/brieefly/model"
)

// BasicProject - a model for basic project info
type BasicProject struct {
	ID          string        `json:"idProject"`
	Name        db.NullString `json:"name"`
	Description db.NullString `json:"description"`
	Type        db.NullString `json:"type"`
	ImageURL    db.NullString `json:"imageUrl"`
}

// Details - a model for agency's details and finished projects
type Details struct {
	Agency         *model.Agency  `json:"agency"`
	AverageOpinion db.NullFloat64 `json:"averageOpinion"`
}
