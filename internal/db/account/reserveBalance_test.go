package account

import (
	"avito_api/internal/db"
	"avito_api/internal/db/model"
	"testing"
)

func TestReserveBalance(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	err = accountDB.CreateAccount(&model.DBCreateAccount{ID: 1234})
	if err != nil {
		t.Fatal("cant create account: ", err)
	}
	acc, err := accountDB.AddBalance(&model.DBAddBalanceInput{ID: 1234, Amount: 100})
	if err != nil {
		t.Fatal("cant add balance: ", err)
	}
	reserve := model.DBReserveBalance{}
	reserve.AccountID = 1234
	reserve.TotalCost = 11.11
	reserve.OperationID = 1212
	reserve.ServiceID = 0
	acc, err = accountDB.ReserveBalance(&reserve)
	if err != nil {
		t.Fatal("cant reserve balance", err)
	}
	if acc.Balance != 88.89 || acc.ReservedBalance != 11.11 {
		t.Fatal("wrong account data after reserve balance")
	}
}
