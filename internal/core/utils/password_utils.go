package utils

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
