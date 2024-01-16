package auth

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(pwd string) (string, error) {
	bytePassword := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
