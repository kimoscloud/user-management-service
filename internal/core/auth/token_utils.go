package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kimoscloud/user-management-service/internal/core/model/dto"
	"time"
)

func GenerateJWT(id string, email string, expirationTime time.Time) (string, error) {
	jwtKey := []byte("your-secret-key") // Usa tu clave secreta aquí

	// Crear las claims
	claims := dto.JWTClaim{
		Email: email,
		ID:    id,
		MapClaims: jwt.MapClaims{
			"iss": "kimos.cloud",
			"sub": id,
			"exp": expirationTime.Unix(),
			"iat": time.Now().Unix(),
		},
	}

	// Crear el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Ejemplo de cómo validar un token
func ValidateToken(signedToken string) (*dto.JWTClaim, error) {
	jwtKey := []byte("your-secret-key") // Usa tu clave secreta aquí

	token, err := jwt.ParseWithClaims(
		signedToken,
		&dto.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*dto.JWTClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
