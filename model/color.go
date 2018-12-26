package model

import "github.com/brieefly/db"

// Color - a color model
type Color struct {
	ID        string        `json:"id_color"`
	HexValue  string        `json:"hex_value"`
	ProjectID db.NullString `json:"project_id"`
}
