package account

import (
	"avito_api/internal/db/model"
	"fmt"
)

func (D *DBAccount) GetBalance(accountID int) (*model.Account, error) {
	if accountID < 0 {
		return nil, fmt.Errorf("account_id must be greater than zero")
	}
	account := model.Account{ID: accountID}
	query := `SELECT balance, reserved_balance FROM account
		WHERE account_id = $1;`
	err := D.db.QueryRow(
		query, accountID,
	).Scan(&account.Balance, &account.ReservedBalance)

	return &account, err
}
