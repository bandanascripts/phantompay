package auth

import (
	"fmt"

	"github.com/bandanascripts/phantompay/pkg/service/app"
	"github.com/bandanascripts/phantompay/pkg/service/app/structs"
	"github.com/bandanascripts/phantompay/pkg/utils"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	app.Connect()
	Db = app.GormDb()
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func FetchLoginCred(username string) ([]byte, string, error) {

	var loginCred structs.PhantomPayUserData

	if err := Db.Select("user_id", "password").Where("username = ?", username).First(&loginCred).Error; err != nil {
		return nil, "", fmt.Errorf("failed to fetch user cred from database : %w", err)
	}

	return []byte(loginCred.Password), loginCred.UserId, nil
}

func Login(u LoginRequest) (string, string, error) {

	hashPassword, userId, err := FetchLoginCred(u.Username)

	if err != nil {
		return "", "", err
	}

	if err = utils.CompareHash(hashPassword, u.Password); err != nil {
		return "", "", err
	}

	return "login successful", userId, nil
}
