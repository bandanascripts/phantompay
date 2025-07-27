package auth

import (
	"fmt"

	"github.com/bandanascripts/phantompay/pkg/service/app/structs"
	"github.com/bandanascripts/phantompay/pkg/utils"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func InsertUser(u structs.PhantomPayUserData) error {

	if err := Db.Create(&u).Error; err != nil {
		return fmt.Errorf("failed to create user data in database : %w", err)
	}

	return nil
}

func Register(u RegisterRequest) (string, error) {

	hashPassword, err := utils.CreateHash(u.Password)

	if err != nil {
		return "", err
	}

	userId, err := utils.NewId()

	if err != nil {
		return "", err
	}

	var userData = structs.PhantomPayUserData{UserId: userId, Username: u.Username, Email: u.Email, Password: string(hashPassword)}

	if err := InsertUser(userData); err != nil {
		return "", err
	}

	return "registeration successful", nil
}
