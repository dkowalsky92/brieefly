package body

import (
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/model"
)

// ProjectDetailsBundle - 
type ProjectDetailsBundle struct {
	ProjectDetails ProjectDetails 			`json:"project"`
	Opinions 		 []model.Opinion 		`json:"opinions"`
	Colors           []model.Color          `json:"colors"`
	Features         []model.Feature        `json:"features"`
	CustomFeatures   []model.CustomFeature  `json:"customFeatures"`
	TargetGroups     []model.TargetGroup    `json:"targetGroups"`
	VisualIdentities []model.VisualIdentity `json:"visualIdentities"`
	SimilarProjects  []model.SimilarProject `json:"similarProjects"`
	Bidder  	  *ProjectBiddingAgency `json:"agency"`
	CMS 	 	  *model.CMS 			`json:"cms"`
	Status  	  *model.ProjectStatus  `json:"status"`
}

// ProjectDetails - project details model
type ProjectDetails struct {
	ProjectID        string                 `json:"idProject"`
	Name             string                 `json:"name"`
	Type             string                 `json:"type"`
	Description      string                 `json:"description"`
	DateCreated      db.NullTime            `json:"dateCreated"`
	DateDeadline     db.NullTime            `json:"dateDeadline"`
	AverageOpinion   float64                `json:"averageOpinion"`
}

// ProjectBiddingAgency - 
type ProjectBiddingAgency struct {
	ID 				string 			`json:"idCompany"`
	Name       		string         	`json:"name"`
	ImageURL    	db.NullString	`json:"imageUrl"`
	AgencyCode      db.NullString 	`json:"agencyCode"`
}
