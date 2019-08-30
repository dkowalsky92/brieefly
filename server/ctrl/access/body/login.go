package body

// LoginInfo - user login info
type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserInfo -
type UserInfo struct {
	Token  string `json:"token"`
	UserID string `json:"idUser"`
	Role   string `json:"role"`
}

// ResetPassword -
type ResetPassword struct {
	Email       string `json:"email"`
	NewPassword string `json:"password"`
}
