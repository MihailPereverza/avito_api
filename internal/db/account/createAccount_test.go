package account

import (
	"avito_api/internal/db"
	"avito_api/internal/db/model"
	"testing"
)

func TestCreateAccountValid(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	err = accountDB.CreateAccount(&model.DBCreateAccount{ID: 1234})
	if err != nil {
		t.Fatal("cant create account: ", err)
	}
	var accountBalance, accountReserveBalance float32 = 0, 0
	db.QueryRow(`SELECT balance, reserved_balance FROM account
		WHERE account_id = 1234`).Scan(&accountBalance, &accountReserveBalance)
	if accountBalance != 0 || accountReserveBalance != 0 {
		t.Fatalf("invalid base balance info: balance=%f, reserveBalance=%f", accountBalance, accountReserveBalance)
	}
}

func TestCreateAccountInvalidData(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	err = accountDB.CreateAccount(&model.DBCreateAccount{ID: -1})
	if err.Error() != `cant create account with ID < 1` {
		t.Fatal("ot return error on incorrect account_id", err)
	}
	var countRow int
	err = db.QueryRow(`SELECT COUNT(account_id) FROM account`).Scan(&countRow)

	// всегда есть аккаунт компании
	if countRow != 1 {
		t.Fatal("create invalid account")
	}
}
