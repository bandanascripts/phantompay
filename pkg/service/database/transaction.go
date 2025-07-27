package database

import (
	"fmt"
	"time"

	"github.com/bandanascripts/phantompay/pkg/service/app/structs"
	"github.com/bandanascripts/phantompay/pkg/utils"
	"gorm.io/gorm"
)

type TransactionRequest struct {
	UserId     string
	RecieverId string  `json:"recieverid"`
	Amount     float64 `json:"amount"`
}

func StoreTransaction(tx *gorm.DB, transaction structs.PhantomPayTransaction) error {

	if err := tx.Create(&transaction).Error; err != nil {
		return fmt.Errorf("failed to store transaction details in database : %w", err)
	}

	return nil
}

func FetchWallets(tx *gorm.DB, transaction TransactionRequest) (*structs.PhantomPayWallet, *structs.PhantomPayWallet, error) {

	var userWallet, receiverWallet structs.PhantomPayWallet

	if err := tx.Where("user_id = ?", transaction.UserId).First(&userWallet).Error; err != nil {
		return nil, nil, fmt.Errorf("failed to get user and wallet : %w", err)
	}

	if err := tx.Where("user_id = ?", transaction.RecieverId).First(&receiverWallet).Error; err != nil {
		return nil, nil, fmt.Errorf("failed to get receiver wallet : %w", err)
	}

	return &userWallet, &receiverWallet, nil
}

func UpdateWallets(tx *gorm.DB, userWallet, receiverWallet *structs.PhantomPayWallet) error {

	if err := tx.Save(userWallet).Error; err != nil {
		return fmt.Errorf("failed to update user wallet : %w", err)
	}

	if err := tx.Save(receiverWallet).Error; err != nil {
		return fmt.Errorf("failed to update receiver wallet : %w", err)
	}

	return nil
}

func Transaction(transaction TransactionRequest) (string, error) {

	var tx = Db.Begin()

	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction : %w", tx.Error)
	}

	userWallet, receiverWallet, err := FetchWallets(tx, transaction)

	if err != nil {
		tx.Rollback()
		return "", err
	}

	userWallet.Balance -= transaction.Amount
	receiverWallet.Balance += transaction.Amount

	if err := UpdateWallets(tx, userWallet, receiverWallet); err != nil {
		tx.Rollback()
		return "", err
	}

	transactionId, err := utils.NewId()

	if err != nil {
		return "", err
	}

	var transactionDetails = structs.PhantomPayTransaction{ID: transactionId, SenderID: userWallet.UserID, ReceiverID: receiverWallet.UserID, Amount: transaction.Amount, CreatedAt: time.Now()}

	if err := StoreTransaction(tx, transactionDetails); err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction : %w", err)
	}

	return "transaction successful", nil
}
