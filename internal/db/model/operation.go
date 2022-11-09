package model

import (
	"avito_api/internal/handler/inputModel/account"
	"time"
)

type Operation struct {
	ID         int     `json:"operation_id"`
	Account    Account `json:"account"`
	Status     Status  `json:"status"`
	Service    `json:"service"`
	Count      int     `json:"count"`
	TotalCost  float32 `json:"total_cost"`
	CreateTime time.Time
}

type DBBaseOperationInfo struct {
	account.ServiceAccountInput
}

type OperationInfo struct {
	ID         int       `json:"operation_id"`
	AccountID  int       `json:"account_id"`
	ServiceID  int       `json:"service_id"`
	StatusID   int       `json:"status_id"`
	TotalCost  float32   `json:"total_cost"`
	CreateTime time.Time `json:"create_time"`
}

type OperationReport struct {
	ID           int       `json:"operation_id"`
	ServiceTitle string    `json:"service_title"`
	AccountID    int       `json:"account_id"`
	TotalCost    float32   `json:"total_cost"`
	CreateTime   time.Time `json:"create_time"`
}
