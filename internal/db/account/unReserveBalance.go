package account

import (
	"avito_api/internal/db/model"
	"context"
	"database/sql"
	"fmt"
)

func (D *DBAccount) UnReserveBalance(baseInfo *model.DBBaseOperationInfo) (*model.Account, error) {
	operationInfo, err := D.getOperationInfo(baseInfo)
	if err != nil {
		return nil, err
	}
	if operationInfo.StatusID != 0 {
		return nil, fmt.Errorf("DBAccount.UnReserveBalance.getOperation cant unreserve operation")
	}

	tx, err := D.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("db.BeginTx %w", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	account, err := updateBalance(tx, operationInfo)
	if err != nil {
		return nil, err
	}
	err = cancelOperation(tx, operationInfo)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("UnReserveBalance.tx.Commit %w", err)
	}
	return account, nil
}

func updateBalance(tx *sql.Tx, operationInfo *model.OperationInfo) (*model.Account, error) {
	account := model.Account{ID: operationInfo.AccountID}
	queryAccount := `UPDATE account
    	SET balance = balance + $1, reserved_balance = reserved_balance - $1
		WHERE account_id = $2
		RETURNING balance, reserved_balance;`
	err := tx.QueryRow(queryAccount, operationInfo.TotalCost, operationInfo.AccountID).Scan(&account.Balance, &account.ReservedBalance)
	if err != nil {
		return nil, fmt.Errorf("UnReserveBalance.updateAccountBalance %w", err)
	}
	return &account, nil
}

func cancelOperation(tx *sql.Tx, operationInfo *model.OperationInfo) error {
	queryOperation := `UPDATE account_operation SET status_id = 2 WHERE 
		account_id = $1 AND status_id = 0 AND service_id = $2 
		AND operation_id = $3 AND total_cost = $4`
	res, err := tx.Exec(queryOperation,
		operationInfo.AccountID, operationInfo.ServiceID, operationInfo.ID, operationInfo.TotalCost,
	)
	fmt.Println(operationInfo)
	if err != nil {
		return fmt.Errorf("UnReserveBalance.cancelOperation %w", err)
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		fmt.Println(affected)
		fmt.Println(res)
		return fmt.Errorf("UnReserveBalance.cancelOperation not affected")
	}
	return nil
}
