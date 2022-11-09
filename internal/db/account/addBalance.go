package account

import "avito_api/internal/db/model"

func (D *DBAccount) AddBalance(accountInput *model.DBAddBalanceInput) (*model.Account, error) {
	account := model.Account{ID: accountInput.ID}
	query := `UPDATE account
    	SET balance = balance + $1
		WHERE account_id = $2
		RETURNING balance, reserved_balance`
	err := D.db.QueryRow(
		query, accountInput.Amount, accountInput.ID,
	).Scan(&account.Balance, &account.ReservedBalance)
	return &account, err
}
