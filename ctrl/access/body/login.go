package body

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// JWTClaims - claims to be used inside JWT token
type JWTClaims struct {
	UserID string `json:"idUser"`
	jwt.StandardClaims
}

// Token - authorization token
type Token struct {
	Token string `json:"token"`
}

// Info - user login info
type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
