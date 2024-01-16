package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.MapClaims
}
