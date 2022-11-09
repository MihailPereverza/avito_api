package model

import (
	"avito_api/internal/handler/inputModel/account"
	"time"
)

type Account struct {
	ID              int     `json:"account_id"`
	Balance         float32 `json:"balance"`
	ReservedBalance float32 `json:"reserved_balance"`
}

type DBCreateAccount struct {
	ID int `json:"account_id"`
}

type DBAddBalanceInput struct {
	ID     int     `json:"account_id"`
	Amount float32 `json:"amount"`
}

type DBGetStatistic struct {
	AccountID       int
	OrderBy         string
	OrderDirection  int
	Direction       int
	LastOperationID int
	Count           int
}

type DBGetStatisticOutput struct {
	OperationID        int
	StatusTitle        string
	ServiceTitle       string
	ServiceDescription string
	TotalCost          float32
	CreateTime         time.Time
}

type DBReserveBalance struct {
	account.ServiceAccountInput
}

type DBTransferInfo struct {
	account.TransferInfo
}
