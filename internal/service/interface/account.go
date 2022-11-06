package service_interface

import "avito_api/internal/db/model"

type AccountServiceInterface interface {
	CreateAccount(account *model.Account) (*model.Account, error)
	AddBalance(account *model.Account, amount float32) (*model.Account, error)
}
