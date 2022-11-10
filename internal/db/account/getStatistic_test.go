package account

import (
	"avito_api/internal/db"
	"avito_api/internal/db/model"
	"database/sql"
	"testing"
)

func TestGetStatisticValid(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	err = accountDB.CreateAccount(&model.DBCreateAccount{ID: 1234})
	if err != nil {
		t.Fatal("cant create account: ", err)
	}
	err = addTestOperation(db, 1234)
	if err != nil {
		t.Fatal("cant add operation: ", err)
	}

	output, err := accountDB.GetStatistic(&model.DBGetStatistic{
		AccountID:       1234,
		OrderBy:         "date",
		OrderDirection:  1,
		Direction:       1,
		LastOperationID: -1,
		Count:           3,
	})
	if err != nil {
		t.Fatal("cant getStatistic1: ", err)
	}
	if output[0].OperationID != 2 || output[1].OperationID != 3 || output[2].OperationID != 4 {
		t.Fatal("invalid statistic1 operationID")
	}
	if output[0].TotalCost != 11.11 || output[1].TotalCost != 11.09 || output[2].TotalCost != 9.11 {
		t.Fatal("invalid statistic1 totalCost")
	}

	output, err = accountDB.GetStatistic(&model.DBGetStatistic{
		AccountID:       1234,
		OrderBy:         "cost",
		OrderDirection:  1,
		Direction:       1,
		LastOperationID: -1,
		Count:           3,
	})
	if err != nil {
		t.Fatal("cant getStatistic2: ", err)
	}
	if output[0].OperationID != 4 || output[1].OperationID != 3 || output[2].OperationID != 2 {
		t.Fatal("invalid statistic2 operationID")
	}
	if output[0].TotalCost != 9.11 || output[1].TotalCost != 11.09 || output[2].TotalCost != 11.11 {
		t.Fatal("invalid statistic2 totalCost")
	}
}

func TestPrepareCompSymbol(t *testing.T) {
	compSymb := prepareCompSymbol(&model.DBGetStatistic{Direction: 1})
	if compSymb != ">" {
		t.Fatal("error on compSymb: expected >, but got ", compSymb)
	}
	compSymb = prepareCompSymbol(&model.DBGetStatistic{Direction: -1})
	if compSymb != "<" {
		t.Fatal("error on compSymb: expected <, but got ", compSymb)
	}
}

func TestPrepareOrderBy(t *testing.T) {
	res, err := prepareOrderBy("date")
	if err != nil || res != "create_time" {
		t.Fatal("not right answer on date")
	}
	res, err = prepareOrderBy("cost")
	if err != nil || res != "total_cost" {
		t.Fatal("not right answer on cost")
	}
	res, err = prepareOrderBy("abobus")
	if err == nil {
		t.Fatal("not right answer on abobus")
	}
}

func TestCheckExistsUserOperation(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	err = accountDB.CreateAccount(&model.DBCreateAccount{ID: 1234})
	if err != nil {
		t.Fatal("cant create account: ", err)
	}
	err = checkExistsUserOperation(db, 1234)
	if err == nil {
		t.Fatal("found unreal operations: ", err)
	}
	err = addTestOperation(db, 1234)
	if err != nil {
		t.Fatal("cant add operations: ", err)
	}
	err = checkExistsUserOperation(db, 1234)
	if err != nil {
		t.Fatal("not found operations: ", err)
	}
}

func addTestOperation(db *sql.DB, accountID int) error {
	_, err := db.Exec(`INSERT INTO account_operation(operation_id, account_id, status_id, service_id, count, total_cost) 
		VALUES (2, $1, 0, 0, 1, 11.11)`, accountID)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO account_operation(operation_id, account_id, status_id, service_id, count, total_cost) 
		VALUES (3, $1, 0, 0, 1, 11.09)`, accountID)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO account_operation(operation_id, account_id, status_id, service_id, count, total_cost) 
		VALUES (4, $1, 0, 0, 1, 9.11)`, accountID)
	if err != nil {
		return err
	}
	return nil
}
