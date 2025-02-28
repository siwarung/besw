package model

import (
	"github.com/golang-jwt/jwt/v4"
)

// UserClaims hanya menyimpan data yang relevan untuk JWT
type UserClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

// JWTClaims sekarang menggunakan UserClaims
type JWTClaims struct {
	jwt.RegisteredClaims
	User UserClaims `json:"user"`
}
