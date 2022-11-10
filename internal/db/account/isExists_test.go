package account

import (
	"avito_api/internal/db"
	"avito_api/internal/db/model"
	"testing"
)

func TestIsExists(t *testing.T) {
	db, err := db.InitTestDB()
	if err != nil {
		t.Fatal(err)
	}

	accountDB := NewDBAccount(db)
	res, err := accountDB.IsExists(1234)
	if err != nil || res == true {
		t.Fatal("error existence account")
	}
	err = accountDB.CreateAccount(&model.DBCreateAccount{ID: 1234})
	if err != nil {
		t.Fatal("cant create account")
	}
	res, err = accountDB.IsExists(1234)
	if err != nil || res == false {
		t.Fatal("error existence account")
	}
}
