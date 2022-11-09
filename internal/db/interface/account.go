package db_interface

import "avito_api/internal/db/model"

type AccountDB interface {
	CreateAccount(account *model.DBCreateAccount) error
	AddBalance(account *model.DBAddBalanceInput) (*model.Account, error)
	GetBalance(accountID int) (*model.Account, error)
	GetStatistic(statistic *model.DBGetStatistic) ([]model.DBGetStatisticOutput, error)
	ReserveBalance(operation *model.DBReserveBalance) (*model.Account, error)
	UnReserveBalance(operation *model.DBBaseOperationInfo) (*model.Account, error)
	IsExists(accountID int) (bool, error)
	Transfer(info *model.DBTransferInfo) (*model.Account, *model.Account, error)
}
