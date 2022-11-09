package service_interface

import (
	"avito_api/internal/db/model"
	"avito_api/internal/handler/inputModel/account"
	account2 "avito_api/internal/handler/outputModel/account"
)

type AccountServiceInterface interface {
	CreateAccount(account *model.DBCreateAccount) error
	AddBalance(account *account.AddBalanceInput) (*account2.AddBalanceOutput, error)
	GetBalance(accountID int) (*model.Account, error)
	GetStatistic(input *account.GetStatisticInput) ([]account2.GetStatisticOutput, error)
	ReserveBalance(account *account.ServiceAccountInput) (*model.Account, error)
	UnReserveBalance(operation *account.ServiceAccountInput) (*model.Account, error)
	ApproveDebiting(operation *account.ServiceAccountInput) (*account2.ApproveDebitingOutput, error)
	Transfer(info *account.TransferInfo) (*account2.TransferOutput, error)
}
