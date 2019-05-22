package model

import "github.com/dkowalsky/brieefly/db"

// Company - a company struct
type Company struct {
	ID               string        `json:"idCompany"`
	Email            string        `json:"email"`
	Name             db.NullString `json:"name"`
	Phone            db.NullString `json:"phone"`
	Address          db.NullString `json:"address"`
	NameURL          db.NullString `json:"nameUrl"`
	WebsiteURL       db.NullString `json:"websiteUrl"`
	ImageURL         db.NullString `json:"imageUrl"`
	Description      db.NullString `json:"description"`
	DateLastModified db.NullTime   `json:"dateLastModified"`
	DateCreated      db.NullTime   `json:"dateCreated"`
}

// -- Table: Company
// CREATE TABLE Company (
//     id_company int NOT NULL AUTO_INCREMENT,
//     name varchar(100) NOT NULL,
//     website_url varchar(300) NULL,
//     phone varchar(20) NULL,
//     email varchar(75) NOT NULL,
//     address varchar(300) NULL,
//     description varchar(300) NULL,
//     image_url varchar(200) NULL,
//     date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     date_last_modified timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
//     CONSTRAINT Company_pk PRIMARY KEY (id_company)
// );
