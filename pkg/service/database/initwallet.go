package database

import (
	"fmt"

	"github.com/bandanascripts/phantompay/pkg/service/app/structs"
	"github.com/bandanascripts/phantompay/pkg/utils"
)

func CreateWallet(wallet structs.PhantomPayWallet) error {

	if err := Db.Create(&wallet).Error; err != nil {
		return fmt.Errorf("failed to create user wallet inside database")
	}

	return nil
}

func InitWallet(userId string) (string, string, error) {

	walletId, err := utils.NewId()

	if err != nil {
		return "", "", err
	}

	var wallet = structs.PhantomPayWallet{UserID: userId, WalletID: walletId}

	if err := CreateWallet(wallet); err != nil {
		return "", "", err
	}

	return "wallet created", walletId, nil
}
