package service

import (
	"avito_api/internal/db/interface"
	"avito_api/internal/db/model"
)

type AccountService struct {
	accountDB db_interface.AccountDB
}

func NewAccountService(dbAccount db_interface.AccountDB) *AccountService {
	return &AccountService{
		dbAccount,
	}
}

func (us *AccountService) CreateAccount(account *model.Account) (*model.Account, error) {
	return us.accountDB.CreateAccount(account)
}

func (us *AccountService) AddBalance(account *model.Account, amount float32) (*model.Account, error) {
	return us.accountDB.AddBalance(account, amount)
}
