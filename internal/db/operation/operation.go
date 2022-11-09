package operation

import (
	"avito_api/internal/db/interface"
	"avito_api/internal/db/model"
	"context"
	"database/sql"
	"fmt"
)

type DBOperation struct {
	db *sql.DB
}

func NewDBOperation(db *sql.DB) db_interface.OperationDB {
	return &DBOperation{db}
}

func (D *DBOperation) CreateOperation(operation *model.Operation) error {
	query := `INSERT INTO account_operation(
		account_id, status_id, service_id, count, total_cost) 
    	VALUES ($1, $2, $3, $4, $5) 
    	RETURNING operation_id;`
	err := D.db.QueryRow(
		query, operation.Account.ID, operation.Status.ID, operation.Service.ID, operation.Count, operation.TotalCost,
	).Scan(&operation.ID)
	return err
}

func (D *DBOperation) GetDB() *sql.DB {
	return D.db
}

func (D *DBOperation) ApproveDebiting(operation *model.DBBaseOperationInfo) (companyBalance float32, err error) {
	tx, err := D.db.BeginTx(context.Background(), nil)
	if err != nil {
		return -1, fmt.Errorf("db.BeginTx %w", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	query := `UPDATE account_operation
	SET status_id = 1
	WHERE operation_id = $1 AND status_id = 0 AND account_id = $2 AND service_id = $3 AND total_cost = $4;`
	res, err := tx.Exec(query, operation.OperationID, operation.AccountID, operation.ServiceID, operation.TotalCost)
	if err != nil {
		return -1, fmt.Errorf("DB.ApproveDebiting %w", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("DB.ApproveDebiting %w", err)
	}

	// так как ни одна строка не изминилась и ошибок нет, то стоит проверить,
	//на то, апрувалась ли уже операция с переданным ID
	if affected == 0 {
		accountID := -1
		query = `SELECT account_id from account_operation
	WHERE operation_id = $1 AND status_id = 1 AND account_id = $2 AND service_id = $3 AND total_cost = $4;`
		err = D.db.QueryRow(query, operation.OperationID, operation.AccountID, operation.ServiceID, operation.TotalCost).Scan(&accountID)
		if accountID == operation.AccountID {
			return -1, fmt.Errorf("DB.ApproveDebiting operation already approved")
		}
		return -1, fmt.Errorf("DB.ApproveDebiting operation not found")
	}

	// записываем деньги на счет компании
	query = `UPDATE account SET balance = balance + $1 WHERE account_id = 0 RETURNING balance;`
	err = tx.QueryRow(query, operation.TotalCost).Scan(&companyBalance)
	fmt.Println(err)
	if err != nil {
		return -1, fmt.Errorf("DB.ApproveDebiting.addCompanyBalance %w", err)
	}
	if err = tx.Commit(); err != nil {
		return -1, fmt.Errorf("ApproveDebiting.tx.Commit %w", err)
	}
	return companyBalance, nil
}
