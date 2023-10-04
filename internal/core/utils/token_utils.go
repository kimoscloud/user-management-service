package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/kimoscloud/user-management-service/internal/core/model/dto"
	"os"
	"time"
)

func GenerateJWT(id string, email string, expirationTime time.Time) (string, error) {
	jwtKey := os.Getenv("JWT_KEY")
	claims := &dto.JWTClaim{
		Email: email,
		ID:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", errors.New("error generating JWT token")
	}
	return tokenString, nil

}
