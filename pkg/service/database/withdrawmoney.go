package database

type WithdrawRequest struct {
	Amount float64 `json:"amount"`
}

func WithdrawMoney(userId string, withdraw WithdrawRequest) (string, error) {

	wallet, err := FetchWallet(userId)

	if err != nil {
		return "", err
	}

	wallet.Balance -= withdraw.Amount

	if err := UpdateWallet(wallet); err != nil {
		return "", err
	}

	return "money withdrawn successfully", nil
}
