package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword gera um hash bcrypt da senha
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("password cannot be empty")
	}

	// bcrypt.DefaultCost = 10 (bom equilíbrio entre segurança e performance)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

// ComparePassword compara uma senha em texto plano com um hash bcrypt
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// IsPasswordValid valida se a senha atende aos requisitos mínimos
func IsPasswordValid(password string) error {
	if len(password) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}
	if len(password) > 72 {
		// bcrypt tem limite de 72 bytes
		return fmt.Errorf("password must be at most 72 characters long")
	}
	return nil
}
