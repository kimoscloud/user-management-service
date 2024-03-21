package utils

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func GeneratePassword(pwd string) (string, error) {
	bytePassword := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
