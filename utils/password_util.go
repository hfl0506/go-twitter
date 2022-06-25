package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func ComparePassword(password string, hashedPassword string) (bool, error) {
	isValid := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if isValid.Error() != "" {
		return false, nil
	}
	return true, nil
}
