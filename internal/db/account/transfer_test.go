package account

import (
	"avito_api/internal/db"
	"avito_api/internal/db/model"
	"testing"
)

func TestTransferBalance(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	err = accountDB.CreateAccount(&model.DBCreateAccount{ID: 1234})
	if err != nil {
		t.Fatal("cant create account: ", err)
	}
	err = accountDB.CreateAccount(&model.DBCreateAccount{ID: 1233})
	if err != nil {
		t.Fatal("cant create account: ", err)
	}
	_, err = accountDB.AddBalance(&model.DBAddBalanceInput{ID: 1234, Amount: 100})
	if err != nil {
		t.Fatal("cant add balance: ", err)
	}
	transferInfo := model.DBTransferInfo{}
	transferInfo.FromAccountID = 1234
	transferInfo.ToAccountID = 1233
	transferInfo.Amount = 5.5
	fromAccount, toAccount, err := accountDB.Transfer(&transferInfo)
	if err != nil {
		t.Fatal("cant transfer")
	}
	if fromAccount.Balance != 94.5 || fromAccount.ReservedBalance != 0 {
		t.Fatal("wrong balance on fromAccount")
	}
	if toAccount.Balance != 5.5 || toAccount.ReservedBalance != 0 {
		t.Fatal("wrong balance on toAccount")
	}
}
