package model

import "github.com/dkowalsky/brieefly/db"

// Company - a company struct
type Company struct {
	ID               string        `json:"idCompany" orm:"id_company"`
	Email            string        `json:"email" orm:"email"`
	Name             db.NullString `json:"name" orm:"name"`
	Phone            db.NullString `json:"phone" orm:"phone"`
	Address          db.NullString `json:"address" orm:"address"`
	NameURL          db.NullString `json:"urlName" orm:"url_name"`
	WebsiteURL       db.NullString `json:"websiteUrl" orm:"website_url"`
	ImageURL         db.NullString `json:"imageUrl" orm:"image_url"`
	Description      db.NullString `json:"description" orm:"description"`
	DateLastModified db.NullTime   `json:"dateLastModified" orm:"date_last_modified"`
	DateCreated      db.NullTime   `json:"dateCreated" orm:"date_created"`
}
