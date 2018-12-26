package model

import "github.com/brieefly/db"

// User - a brieefly user
type User struct {
	ID               string        `json:"idUser"`
	Email            string        `json:"email"`
	Login            db.NullString `json:"login"`
	Name             db.NullString `json:"name"`
	Surname          db.NullString `json:"surname"`
	Phone            db.NullString `json:"phone"`
	WebsiteURL       db.NullString `json:"websiteUrl"`
	ImageURL         db.NullString `json:"imageUrl"`
	Description      db.NullString `json:"description"`
	DateOfBirth      db.NullTime   `json:"dateOfBirth"`
	DateLastLogged   db.NullTime   `json:"dateLastLogged"`
	DateCreated      db.NullTime   `json:"dateCreated"`
	DateLastModified db.NullTime   `json:"dateLastModified"`
}
