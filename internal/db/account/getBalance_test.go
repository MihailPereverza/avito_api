package account

import (
	"avito_api/internal/db"
	"avito_api/internal/db/model"
	"testing"
)

func TestGetBalanceValid(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	err = accountDB.CreateAccount(&model.DBCreateAccount{ID: 1234})
	if err != nil {
		t.Fatal("cant create account: ", err)
	}
	_, err = accountDB.AddBalance(&model.DBAddBalanceInput{ID: 1234, Amount: 10.11})
	if err != nil {
		t.Fatal("cant add balance to account: ", err)
	}
	resGet, err := accountDB.GetBalance(1234)
	if err != nil || resGet.Balance != 10.11 || resGet.ReservedBalance != 0 {
		t.Fatal("cant create add balance to account: ", err)
	}
}

func TestGetBalanceInvalidData(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	err = accountDB.CreateAccount(&model.DBCreateAccount{ID: 1234})
	if err != nil {
		t.Fatal("cant create account: ", err)
	}
	_, err = accountDB.AddBalance(&model.DBAddBalanceInput{ID: 1234, Amount: 11.11})
	if err != nil {
		t.Fatal("cant add balance to account: ", err)
	}
	_, err = accountDB.GetBalance(-1)
	if err == nil {
		t.Fatal("not return error on invalid account_id")
	}
}
