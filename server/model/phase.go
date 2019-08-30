package model

import (
	"time"

	"github.com/dkowalsky/brieefly/db"
)

// Phase - a model for project phase
type Phase struct {
	ID            string        `json:"idPhase" orm:"id_phase"`
	Name          string        `json:"name" orm:"name"`
	IsActive      bool          `json:"isActive" orm:"is_active"`
	Description   db.NullString `json:"description" orm:"description"`
	Value         int64         `json:"value" orm:"value"`
	Progress      int64         `json:"progress" orm:"progress"`
	OrderPosition db.NullInt64  `json:"orderPosition" orm:"order_position"`
	Status        db.NullString `json:"status" orm:"status"`
	DateCreated   time.Time     `json:"dateCreated" orm:"date_created"`
}
