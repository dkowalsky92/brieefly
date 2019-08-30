package body

// Password -
type Password struct {
	Password string `json:"password"`
}

// NewPassword -
func NewPassword(password string) *Password {
	return &Password{
		Password: password,
	}
}
