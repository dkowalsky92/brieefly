package body

import (
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/model"
	"github.com/dkowalsky/brieefly/util"
)

// Body -
type Body struct {
	Project         ProjectBody          `json:"project" orm:"-"`
	VisualIdentity  *VisualIdentityBody   `json:"visualIdentity" orm:"-"`
	Colors          []ColorBody          `json:"colors" orm:"-"`
	Features        []FeatureBody        `json:"features" orm:"-"`
	TargetGroups    []TargetGroupBody    `json:"targetGroups" orm:"-"`
	CustomFeatures  []CustomFeatureBody  `json:"customFeatures" orm:"-"`
	SimilarProjects []SimilarProjectBody `json:"similarProjects" orm:"-"`
}

// ProjectBody -
type ProjectBody struct {
	Name         string        `json:"name" orm:"name"`
	Type         string        `json:"type" orm:"type"`
	Description  string        `json:"description" orm:"description"`
	Language     db.NullString `json:"language" orm:"language"`
	BudgetMin    db.NullInt64  `json:"budgetMin" orm:"budget_min"`
	BudgetMax    db.NullInt64  `json:"budgetMax" orm:"budget_max"`
	SubpageCount db.NullInt64  `json:"subpageCount" orm:"subpage_count"`
	NameURL      db.NullString `json:"urlName" orm:"url_name"`
	DateDeadline db.NullTime   `json:"dateDeadline" orm:"date_deadline"`
	CmsID        db.NullString `json:"idCms" orm:"id_cms"`
}

// Project -
type Project struct {
	model.Project
	CmsID    db.NullString `json:"idCms" orm:"id_cms"`
	StatusID string `json:"idStatus" orm:"id_status"`
}

// NewProject -
func NewProject(body ProjectBody, cmsid db.NullString, status string) Project {
	uuid := util.UUID().String()
	var url string
	if body.NameURL.Valid == false {
		url = db.ParseSlug(body.Name)
	} else {
		url = db.ParseSlug(body.NameURL.String)
	}
	return Project{
		model.Project{
			ID:           uuid,
			Name:         body.Name,
			Type:         body.Type,
			Description:  body.Description,
			Language:     body.Language,
			BudgetMin:    body.BudgetMin,
			BudgetMax:    body.BudgetMax,
			SubpageCount: body.SubpageCount,
			NameURL:      url,
			DateDeadline: body.DateDeadline,
		},
		cmsid,
		status,
	}
}

// ClientProjectBody -
type ClientProjectBody struct {
	UserID    string `json:"idUser" orm:"id_user"`
	ProjectID string `json:"idProject" orm:"id_project"`
}

// ClientProject -
type ClientProject struct {
	ClientProjectBody
}

// NewClientProject -
func NewClientProject(body ClientProjectBody) ClientProject {
	return ClientProject{
		body,
	}
}

// FeatureBody -
type FeatureBody struct {
	ID string `json:"idFeature" orm:"id_feature"`
}

// Feature -
type Feature struct {
	FeatureBody
	ProjectID string `json:"idProject" orm:"id_project"`
}

// NewFeature -
func NewFeature(body FeatureBody, projectid string) Feature {
	return Feature{
		body,
		projectid,
	}
}

// CustomFeatureBody -
type CustomFeatureBody struct {
	Name        string `json:"name" orm:"name"`
	Description string `json:"description" orm:"description"`
}

// CustomFeature -
type CustomFeature struct {
	model.CustomFeature
	ProjectID string `json:"idProject" orm:"id_project"`
}

// NewCustomFeature -
func NewCustomFeature(body CustomFeatureBody, projectid string) CustomFeature {
	uuid := util.UUID().String()
	return CustomFeature{
		model.CustomFeature{
			ID:          uuid,
			Name:        body.Name,
			Description: body.Description,
		},
		projectid,
	}
}

// SimilarProjectBody -
type SimilarProjectBody struct {
	ProjectURL string `json:"url" orm:"project_url"`
}

// SimilarProject -
type SimilarProject struct {
	model.SimilarProject
	ProjectID string `json:"idProject" orm:"id_project"`
}

// NewSimilarProject -
func NewSimilarProject(body SimilarProjectBody, projectid string) SimilarProject {
	uuid := util.UUID().String()
	return SimilarProject{
		model.SimilarProject{
			ID:         uuid,
			ProjectURL: body.ProjectURL,
		},
		projectid,
	}
}

// VisualIdentityBody -
type VisualIdentityBody struct {
	Type string `json:"type" orm:"type"`
}

// VisualIdentity -
type VisualIdentity struct {
	model.VisualIdentity
	ProjectID string `json:"id_project" orm:"id_project"`
}

// NewVisualIdentity -
func NewVisualIdentity(body VisualIdentityBody, projectid string) VisualIdentity {
	uuid := util.UUID().String()
	return VisualIdentity{
		model.VisualIdentity{
			ID:   uuid,
			Type: body.Type,
		},
		projectid,
	}
}

// ColorBody -
type ColorBody struct {
	HexValue string `json:"hexValue" orm:"hex_value"`
}

// Color -
type Color struct {
	model.Color
	ProjectID string `json:"idProject" orm:"id_project"`
}

// NewColor -
func NewColor(body ColorBody, projectid string) Color {
	uuid := util.UUID().String()
	return Color{
		model.Color{
			ID:       uuid,
			HexValue: body.HexValue,
		},
		projectid,
	}
}

// TargetGroupBody -
type TargetGroupBody struct {
	Name        string        `json:"name" orm:"name"`
	Description db.NullString `json:"description" orm:"description"`
	AgeMin      db.NullInt64  `json:"ageMin" orm:"age_min"`
	AgeMax      db.NullInt64  `json:"ageMax" orm:"age_max"`
}

// TargetGroup -
type TargetGroup struct {
	model.TargetGroup
	ProjectID string `json:"idProject" orm:"id_project"`
}

// NewTargetGroup -
func NewTargetGroup(body TargetGroupBody, projectid string) TargetGroup {
	uuid := util.UUID().String()
	return TargetGroup{
		model.TargetGroup{
			ID:          uuid,
			Name:        body.Name,
			Description: body.Description,
			AgeMin:      body.AgeMin,
			AgeMax:      body.AgeMax,
		},
		projectid,
	}
}
