package account

import (
	"avito_api/internal/db/interface"
	"avito_api/internal/db/model"
	"database/sql"
)

type DBAccount struct {
	db *sql.DB
}

func NewDBAccount(db *sql.DB) db_interface.AccountDB {
	return &DBAccount{db}
}

func (D *DBAccount) getOperationInfo(baseInfo *model.DBBaseOperationInfo) (*model.OperationInfo, error) {
	query := `SELECT status_id, total_cost, create_time from account_operation
	WHERE operation_id = $1 AND account_id = $2 AND service_id = $3 ;`

	operation := model.OperationInfo{}
	operation.ID = baseInfo.OperationID
	operation.AccountID = baseInfo.AccountID
	operation.ServiceID = baseInfo.ServiceID

	err := D.db.QueryRow(
		query, baseInfo.OperationID, baseInfo.AccountID, baseInfo.ServiceID,
	).Scan(&operation.StatusID, &operation.TotalCost, &operation.CreateTime)
	return &operation, err
}
