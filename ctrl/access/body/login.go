package body

// Token - authorization token
type Token struct {
	UserID string `json:"idUser"`
	Token  string `json:"token"`
}

// LoginInfo - user login info
type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
