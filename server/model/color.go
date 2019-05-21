package model

// Color - a color model
type Color struct {
	ID       string `json:"idColor" orm:"id_color"`
	HexValue string `json:"hexValue" orm:"hex_value"`
}
