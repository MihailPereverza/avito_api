package account

import (
	"avito_api/internal/service/interface"
	"errors"
)

var balanceGreaterZero = errors.New(`balance must be greater zero`)

var duplicateError = "ERROR: duplicate key value violates unique constraint \"account_user_id_currency_id_key\" (SQLSTATE 23505)"
var notFoundCurrency = `ERROR: insert or update on table "account" violates foreign key constraint "account_currency_id_fkey" (SQLSTATE 23503)`
var notFoundUser = `ERROR: insert or update on table "account" violates foreign key constraint "account_user_id_fkey" (SQLSTATE 23503)`
var notFoundAccountToSet = `sql: no rows in result set`

var userIdRequiredError = errors.New(`user_id is required param`)
var currencyIDRequiredError = errors.New(`currency_id is required param`)
var amountRequiredError = errors.New("amount is required param")

var somethingError = errors.New(`something error in request data`)

type AccountHandler struct {
	accountService service_interface.AccountServiceInterface
}

func NewAccountHandler(accountService service_interface.AccountServiceInterface) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}
