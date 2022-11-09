package operation

import (
	"avito_api/internal/db/model"
	"fmt"
)

func (D *DBOperation) GetApproved() ([]model.OperationReport, error) {
	query := `SELECT 
    	operation_id,
		s.title,
		account_id,  
		total_cost,
		create_time
	FROM 
		account_operation
		inner join operation_status os on os.status_id = account_operation.status_id
		inner join service s on s.service_id = account_operation.service_id
	WHERE
		os.status_id = 1
		AND create_time >= (SELECT CURRENT_DATE - INTERVAL '1 mon');`
	rows, err := D.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("cant")
	}
	var operations []model.OperationReport

	for rows.Next() {
		var operation model.OperationReport
		err = rows.Scan(&operation.ID, &operation.ServiceTitle,
			&operation.AccountID, &operation.TotalCost, &operation.CreateTime,
		)
		if err != nil {
			return operations, err
		}
		operations = append(operations, operation)
	}
	if err = rows.Err(); err != nil {
		return operations, err
	}
	return operations, nil
}
