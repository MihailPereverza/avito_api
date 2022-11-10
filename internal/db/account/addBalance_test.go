package account

import (
	"avito_api/internal/db"
	"avito_api/internal/db/model"
	"testing"
)

func TestAddBalanceValid(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	var accountBalance float32 = 23.23
	_, err = db.Exec(`INSERT INTO account(account_id) VALUES (1234)`)
	if err != nil {
		t.Fatal("cant create account: ", err)
	}
	getBalance, err := accountDB.AddBalance(&model.DBAddBalanceInput{ID: 1234, Amount: accountBalance})
	if err != nil || getBalance.Balance != accountBalance || getBalance.ReservedBalance != 0 {
	}
}

func TestAddBalanceInvalidData(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	var accountBalance float32 = 23.23
	_, err = db.Exec(`INSERT INTO account(account_id) VALUES (1234)`)
	if err != nil {
		t.Fatal("cant create account")
	}

	_, err = accountDB.AddBalance(&model.DBAddBalanceInput{ID: -1, Amount: accountBalance})
	if err.Error() != `sql: no rows in result set` {
		t.Fatal("not return error on incorrect account_id")
	}
	_, err = accountDB.AddBalance(&model.DBAddBalanceInput{ID: 1234, Amount: -10})
	if err.Error() != `ERROR: new row for relation "account" violates check constraint "account_balance_check" (SQLSTATE 23514)` {
		t.Fatal("not return error on incorrect amount")
	}
}
