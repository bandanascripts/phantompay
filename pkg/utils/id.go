package utils

import (
	"fmt"
	"github.com/google/uuid"
)

func NewId() (string, error) {

	var id = uuid.New()

	if id.String() == "" {
		return "", fmt.Errorf("failed to create new id")
	}

	return id.String(), nil
}
