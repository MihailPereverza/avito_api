package account

import (
	"avito_api/internal/db/model"
	"context"
	"fmt"
)

func (D *DBAccount) ReserveBalance(operation *model.DBReserveBalance) (*model.Account, error) {
	tx, err := D.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("db.BeginTx %w", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	account := model.Account{}
	queryAccount := `UPDATE account
    	SET balance = balance - $1, reserved_balance = reserved_balance + $1
		WHERE account_id = $2
		RETURNING balance, reserved_balance;`
	queryOperation := `INSERT INTO account_operation(
		account_id, status_id, service_id, count, total_cost, operation_id) 
    	VALUES ($1, $2, $3, $4, $5, $6) 
    	RETURNING operation_id;`

	err = tx.QueryRow(queryAccount, operation.TotalCost, operation.AccountID).Scan(&account.Balance, &account.ReservedBalance)
	if err != nil {
		return nil, fmt.Errorf("ReserveBalance.updateAccountBalance %w", err)
	}
	err = tx.QueryRow(queryOperation,
		operation.AccountID, 0, operation.ServiceID, 1, operation.TotalCost, operation.OperationID,
	).Scan(&operation.OperationID)
	if err != nil {
		return nil, fmt.Errorf("ReserveBalance.insertOperation %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("ReserveBalance.tx.Commit %w", err)
	}
	return &account, nil
}
