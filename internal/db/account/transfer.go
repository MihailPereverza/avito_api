package account

import (
	"avito_api/internal/db/model"
	"context"
	"fmt"
)

func (D *DBAccount) Transfer(info *model.DBTransferInfo) (*model.Account, *model.Account, error) {
	fromAccount := model.Account{ID: info.FromAccountID}
	toAccount := model.Account{ID: info.ToAccountID}

	tx, err := D.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("db.BeginTx %w", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	queryFrom := `UPDATE account SET balance = balance - $1
		WHERE account_id = $2 RETURNING balance, reserved_balance;`
	queryTo := `UPDATE account SET balance = balance + $1
		WHERE account_id = $2 RETURNING balance, reserved_balance;`
	queryTransferInsert := `INSERT INTO account_transfer(from_account_id, to_account_id, amount) 
		VALUES ($1, $2, $3)`
	err = D.db.QueryRow(queryFrom, info.Amount, info.FromAccountID).Scan(&fromAccount.Balance, &fromAccount.ReservedBalance)
	if err != nil {
		return nil, nil, err
	}
	err = D.db.QueryRow(queryTo, info.Amount, info.ToAccountID).Scan(&toAccount.Balance, &toAccount.ReservedBalance)
	if err != nil {
		return nil, nil, err
	}
	_, err = D.db.Exec(queryTransferInsert, info.FromAccountID, info.ToAccountID, info.Amount)
	if err != nil {
		return nil, nil, err
	}
	return &fromAccount, &toAccount, err
}
