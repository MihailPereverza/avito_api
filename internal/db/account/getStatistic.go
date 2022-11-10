package account

import (
	"avito_api/internal/config"
	"avito_api/internal/db/model"
	"database/sql"
	"fmt"
	"log"
)

func (D *DBAccount) GetStatistic(stat *model.DBGetStatistic) ([]model.DBGetStatisticOutput, error) {
	err := checkExistsUserOperation(D.db, stat.AccountID)
	if err != nil {
		return nil, fmt.Errorf("account has not operations")
	}

	var sqlOrderDirection string
	sqlOrderDirection, stat.Direction, err = prepareSQLOrderDirection(stat.OrderDirection, stat.Direction)
	if err != nil {
		return nil, err
	}
	stat.OrderBy, err = prepareOrderBy(stat.OrderBy)
	if err != nil {
		return nil, err
	}
	compSymbol := prepareCompSymbol(stat)
	rows, err := D.getOperationRows(stat, compSymbol, sqlOrderDirection)
	if err != nil {
		log.Printf("DBAccount.GetStatistic.db.Query handle err: %s", err)
		return nil, fmt.Errorf("cant create statistic")
	}
	operations, err := operateSQLRows(rows)
	if err = rows.Err(); err != nil {
		log.Printf("DBAccount.GetStatistic.rows.Err handle err: %s", err)
		return nil, fmt.Errorf("cant create statistic")
	}
	return operations, nil
}

func checkExistsUserOperation(db *sql.DB, accountID int) error {
	query := `SELECT account_id FROM account_operation WHERE account_id = $1 LIMIT 1`
	err := db.QueryRow(query, accountID).Scan(&accountID)
	return err
}

func prepareCompSymbol(stat *model.DBGetStatistic) string {
	compSymbol := ""
	switch stat.Direction {
	case 1:
		compSymbol = ">"
	case -1:
		compSymbol = "<"
	}
	return compSymbol
}

func prepareOrderBy(orderBy string) (string, error) {
	switch orderBy {
	case "date":
		orderBy = "create_time"
	case "cost":
		orderBy = "total_cost"
	default:
		return "", fmt.Errorf("order_by must be date or cost")
	}
	return orderBy, nil
}

func prepareSQLOrderDirection(orderDirection int, direction int) (string, int, error) {
	sqlOrderDirection := ""
	switch orderDirection {
	case 0:
		sqlOrderDirection = "DESC"
		direction *= -1
	case 1:
		sqlOrderDirection = "ASC"
	default:
		return "", 0, fmt.Errorf("orderDirection must be 1 or 0")
	}
	return sqlOrderDirection, direction, nil
}

func (D *DBAccount) getOperationRows(stat *model.DBGetStatistic, compSymbol string, sqlOrderDirection string) (*sql.Rows, error) {
	paginationSQL := ""
	if stat.LastOperationID != -1 {
		paginationSQL = fmt.Sprintf(
			`AND (%s, operation_id) %s ((SELECT %s from account_operation where operation_id = %d), %d)`,
			stat.OrderBy, compSymbol, stat.OrderBy, stat.LastOperationID, stat.LastOperationID,
		)
	}

	query := fmt.Sprintf(`SELECT operation_id, os.title, s.title, s.description, total_cost, create_time FROM
		account_operation
		inner join operation_status os on os.status_id = account_operation.status_id
		inner join service s on s.service_id = account_operation.service_id
	WHERE
	    account_id = $1 
	    %s
	ORDER BY %s %s, operation_id %s
	LIMIT $2 OFFSET $3;`, paginationSQL, stat.OrderBy, sqlOrderDirection, sqlOrderDirection)
	rows, err := D.db.Query(query, stat.AccountID, stat.Count, stat.Count-config.GetAppConfig().AccountStatisticPageSize)
	return rows, err
}

func operateSQLRows(rows *sql.Rows) ([]model.DBGetStatisticOutput, error) {
	var err error
	var operations []model.DBGetStatisticOutput
	for rows.Next() {
		var operation model.DBGetStatisticOutput
		err = rows.Scan(&operation.OperationID, &operation.StatusTitle,
			&operation.ServiceTitle, &operation.ServiceDescription,
			&operation.TotalCost, &operation.CreateTime,
		)
		if err != nil {
			log.Printf("DBAccount.GetStatistic.rows.Next handle err: %s", err)
			return nil, fmt.Errorf("cant create statistic")
		}
		operations = append(operations, operation)
	}
	return operations, nil
}
