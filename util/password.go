package util

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

func CheckPassword(password string, hashedPassword string) error {
<<<<<<< HEAD
	
=======
>>>>>>> 51b03b153c0863bf9672ec6395bf295e47c685ac
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
