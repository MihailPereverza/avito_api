package account

import (
	"avito_api/internal/db"
	"avito_api/internal/db/model"
	"testing"
)

func TestUnReserveBalance(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	err = accountDB.CreateAccount(&model.DBCreateAccount{ID: 1234})
	if err != nil {
		t.Fatal("cant create account: ", err)
	}
	_, err = accountDB.AddBalance(&model.DBAddBalanceInput{ID: 1234, Amount: 100})
	if err != nil {
		t.Fatal("cant add balance: ", err)
	}
	reserve := model.DBReserveBalance{}
	reserve.AccountID = 1234
	reserve.TotalCost = 11.5
	reserve.OperationID = 1212
	reserve.ServiceID = 0
	_, err = accountDB.ReserveBalance(&reserve)
	if err != nil {
		t.Fatal("cant reserve balance", err)
	}

	operation := model.DBBaseOperationInfo{}
	operation.OperationID = 1212
	operation.TotalCost = 11.5
	operation.AccountID = 1234
	operation.ServiceID = 0
	res, err := accountDB.UnReserveBalance(&operation)
	if err != nil {
		t.Fatal("cant unreserve: ", err)
	}
	if res.Balance != 100 || res.ReservedBalance != 0 {
		t.Fatal("wrong balance after UnReserveBalance")
	}
}
