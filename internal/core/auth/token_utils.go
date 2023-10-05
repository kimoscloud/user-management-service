package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/kimoscloud/user-management-service/internal/core/model/dto"
	"os"
	"time"
)

func GenerateJWT(id string, email string, expirationTime time.Time) (
	string,
	error,
) {
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
		return "", errors.New("errors generating JWT token")
	}
	return tokenString, nil
}

func ValidateToken(signedToken string) (claims *dto.JWTClaim, err error) {
	jwtKey := os.Getenv("JWT_KEY")
	token, err := jwt.ParseWithClaims(
		signedToken,
		&dto.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*dto.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if (claims.ExpiresAt) < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
