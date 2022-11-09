package account

type TransferInfo struct {
	FromAccountID int     `json:"from_account_id"`
	ToAccountID   int     `json:"to_account_id"`
	Amount        float32 `json:"amount"`
}
