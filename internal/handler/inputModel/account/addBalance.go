package account

type AddBalanceInput struct {
	ID     int     `json:"account_id"`
	Amount float32 `json:"amount"`
}
