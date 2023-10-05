package dto

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTClaim struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
}
