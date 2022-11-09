package db_interface

import (
	"avito_api/internal/db/model"
)

type OperationDB interface {
	CreateOperation(operation *model.Operation) error
	ApproveDebiting(operation *model.DBBaseOperationInfo) (companyBalance float32, err error)
	GetApproved() ([]model.OperationReport, error)
}
