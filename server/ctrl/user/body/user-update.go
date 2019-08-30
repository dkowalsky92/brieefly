package body

import "github.com/dkowalsky/brieefly/db"

// UserUpdate -
type UserUpdate struct {
	Email            db.NullString `json:"email" orm:"email"`
	Login            db.NullString `json:"login" orm:"login"`
	Name             db.NullString `json:"name" orm:"name"`
	Surname          db.NullString `json:"surname" orm:"surname"`
	Phone            db.NullString `json:"phone" orm:"phone"`
	WebsiteURL       db.NullString `json:"websiteUrl" orm:"website_url"`
	ImageURL         db.NullString `json:"imageUrl" orm:"image_url"`
	Description      db.NullString `json:"description" orm:"description"`
	DateOfBirth      db.NullTime   `json:"dateOfBirth" orm:"date_of_birth"`
}
