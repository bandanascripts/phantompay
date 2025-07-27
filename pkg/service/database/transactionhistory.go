package database

import (
	"fmt"

	"github.com/bandanascripts/phantompay/pkg/service/app/structs"
)

func TransactionHistory(userId string) ([]structs.PhantomPayTransaction, error) {

	var transactions []structs.PhantomPayTransaction

	if err := Db.Where("sender_id = ? or receiver_id = ?", userId, userId).Order("created_at DESC").Find(&transactions).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch transactions details for user : %w", err)
	}

	return transactions, nil
}
