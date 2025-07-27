package database

import (
	"fmt"
	"github.com/bandanascripts/phantompay/pkg/service/app"
	"github.com/bandanascripts/phantompay/pkg/service/app/structs"
	"gorm.io/gorm"
)

type DepositRequest struct {
	Amount float64 `json:"amount"`
}

var Db *gorm.DB

func init() {
	app.Connect() ; Db = app.GormDb()
}

func FetchWallet(userId string) (*structs.PhantomPayWallet, error) {

	var wallet structs.PhantomPayWallet

	if err := Db.Where("user_id = ?", userId).First(&wallet).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch wallet details from database : %w", err)
	}

	return &wallet, nil
}

func UpdateWallet(wallet *structs.PhantomPayWallet) error {

	if err := Db.Save(wallet).Error; err != nil {
		return fmt.Errorf("failed to update wallet details : %w", err)
	}

	return nil
}

func DepositMoney(userId string, deposit DepositRequest) (string, error) {

	wallet, err := FetchWallet(userId)

	if err != nil {
		return "", err
	}

	wallet.Balance += deposit.Amount

	if err := UpdateWallet(wallet); err != nil {
		return "", err
	}

	return "money deposited successfully", nil
}
