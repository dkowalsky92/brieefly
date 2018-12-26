package model

import (
	"time"

	"github.com/brieefly/db"
)

// Project - a project model
type Project struct {
	ID               string         `json:"idProject"`
	Name             string         `json:"name"`
	Type             string         `json:"type"`
	Description      string         `json:"description"`
	OverallProgress  int64          `json:"overallProgress"`
	Language         db.NullString  `json:"language"`
	BudgetMin        db.NullInt64   `json:"budgetMin"`
	BudgetMax        db.NullInt64   `json:"budgetMax"`
	SubpageCount     db.NullInt64   `json:"subpageCount"`
	ImageURL         db.NullString  `json:"imageUrl"`
	DateDeadline     db.NullTime    `json:"dateDeadline"`
	DateCreated      time.Time      `json:"dateCreated"`
	DateLastModified db.NullTime    `json:"dateLastModified"`
	Cms              *CMS           `json:"cms"`
	Status           *ProjectStatus `json:"projectStatus"`
}

// id_project int NOT NULL AUTO_INCREMENT,
//     name varchar(70) NOT NULL,
//     type varchar(40) NOT NULL,
//     description varchar(500) NOT NULL,
//     language varchar(50) NULL,
//     budget_min int NULL,
//     budget_max int NULL,
//     subpage_count int NULL,
//     overall_progress int NOT NULL DEFAULT 0,
//     image_url varchar(100) NULL,
//     date_deadline date NULL,
//     date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     date_last_modified timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
//     id_status int NOT NULL,
//     id_cms int NULL,
//     CONSTRAINT Project_pk PRIMARY KEY (id_project)
