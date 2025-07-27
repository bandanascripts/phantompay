package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func CreateHash(password string) ([]byte, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("failed to generate hash from password : %w", err)
	}

	return hashPassword, nil
}

func CompareHash(hashPassword []byte , password string) error {

	if err := bcrypt.CompareHashAndPassword(hashPassword , []byte(password)) ; err != nil {
		return fmt.Errorf("failed to compare hash and password : %w" , err)
	}

	return nil
}