package dto

import "github.com/golang-jwt/jwt"

type Claims struct {
	UserID   uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
