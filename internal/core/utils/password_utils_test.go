package utils

import (
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	pwd := "testpassword123"
	hashedPwd, err := GeneratePassword(pwd)

	if err != nil {
		t.Errorf("Error al generar la contraseña: %s", err)
	}

	if len(hashedPwd) == 0 {
		t.Errorf("La contraseña generada está vacía")
	}
}

func TestComparePasswords(t *testing.T) {
	pwd := "testpassword123"
	hashedPwd, _ := GeneratePassword(pwd)

	testCases := []struct {
		name     string
		hash     string
		password string
		expected bool
	}{
		{"Correcta", hashedPwd, pwd, true},
		{"Incorrecta", hashedPwd, "wrongpassword", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ComparePasswords(tc.hash, tc.password)
			if result != tc.expected {
				t.Errorf("Se esperaba %v, se obtuvo %v", tc.expected, result)
			}
		})
	}
}
