package model

import "github.com/brieefly/db"

// User - a brieefly user
type User struct {
	ID               string        `json:"id_user"`
	Email            string        `json:"email"`
	Login            db.NullString `json:"login"`
	Name             db.NullString `json:"name"`
	Surname          db.NullString `json:"surname"`
	Phone            db.NullString `json:"phone"`
	WebsiteURL       db.NullString `json:"website_url"`
	ImageURL         db.NullString `json:"image_url"`
	Description      db.NullString `json:"description"`
	DateOfBirth      db.NullTime   `json:"date_of_birth"`
	DateLastLogged   db.NullTime   `json:"date_last_logged"`
	DateCreated      db.NullTime   `json:"date_created"`
	DateLastModified db.NullTime   `json:"date_last_modified"`
}
