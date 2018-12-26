package model

import (
	"time"

	"github.com/brieefly/db"
)

// ProjectPhase - a model for project phase
type ProjectPhase struct {
	ID            string        `json:"idProjectPhase"`
	Value         int64         `json:"value"`
	Progress      int64         `json:"progress"`
	OrderPosition db.NullInt64  `json:"orderPosition"`
	Status        db.NullString `json:"status"`
	DateCreated   time.Time     `json:"dateCreated"`
	PhaseID       string        `json:"idPhase"`
}

// -- Table: Project_phase
// CREATE TABLE Project_phase (
//     id_project_phase int NOT NULL AUTO_INCREMENT,
//     value int NOT NULL DEFAULT 1,
//     progress int NOT NULL DEFAULT 0,
//     order_position int NULL,
//     status varchar(20) NOT NULL,
//     date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     id_phase int NOT NULL,
//     id_project int NOT NULL,
//     CONSTRAINT Project_phase_pk PRIMARY KEY (id_project_phase)
// );
