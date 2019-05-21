package body

// RegisterInfo - user registration information necessary for creating an account
type RegisterInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// BasicUser - a user object returned after registration
type BasicUser struct {
	UserID string `json:"idUser"`
	Email  string `json:"email"`
}
