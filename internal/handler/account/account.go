package account

import (
	"avito_api/internal/handler/inputModel/account"
	handler_interface "avito_api/internal/handler/interface"
	"avito_api/internal/service/interface"
	"errors"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"log"
	"net/http"
)

var balanceGreaterZero = errors.New(`balance must be greater zero`)
var enoughMoney = `ERROR: new row for relation "account" violates check constraint "account_balance_check" (SQLSTATE 23514)`
var insufficientFunds = `ReserveBalance.updateAccountBalance ERROR: new row for relation "account" violates check constraint "account_balance_check" (SQLSTATE 23514)`
var notAffectedRows = `DB.ApproveDebiting operation not found`

var notFoundAccountToSet = `sql: no rows in result set`
var notFoundAccountToUpdate = `ReserveBalance.updateAccountBalance sql: no rows in result set`
var notFoundOperation = `ReserveBalance.insertOperation ERROR: insert or update on table "account_operation" violates foreign key constraint "account_operation_service_id_fkey" (SQLSTATE 23503)`
var notFoundAccountOperation = `account has not operations`

var operationAlreadyExists = `ReserveBalance.insertOperation ERROR: duplicate key value violates unique constraint "account_operation_pkey" (SQLSTATE 23505)`
var operationAlreadyApproved = `DB.ApproveDebiting operation already approved`
var operationCantUnReserve = `DBAccount.UnReserveBalance.getOperation cant unreserve operation`

var accountIdRequiredError = errors.New(`account_id is required param`)
var operationIdRequiredError = errors.New(`operation_id is required param`)
var amountRequiredError = errors.New("amount is required param")

var somethingError = errors.New(`something error in request data`)

type AccountHandler struct {
	accountService    service_interface.AccountServiceInterface
	defaultCurrencyID int
}

func NewAccountHandler(
	accountService service_interface.AccountServiceInterface,
) handler_interface.AccountHandlerInterface {
	return &AccountHandler{
		accountService: accountService,
	}
}

func jsonError(w http.ResponseWriter, err error) {
	errorJSON := gabs.New()
	_, err = errorJSON.Set(err.Error(), "error")
	if err != nil {
		errorJSON := gabs.New()
		_, err = errorJSON.Set("something error", "error")
		http.Error(w, errorJSON.String(), http.StatusBadRequest)
		return
	}
	http.Error(w, errorJSON.String(), http.StatusBadRequest)
}

func parseJSONServiceAccount(r *http.Request) (*account.ServiceAccountInput, error) {
	jsonParsed, err := gabs.ParseJSONBuffer(r.Body)
	if err != nil {
		log.Printf("handle error: %s| in parseJSONAddBalance ParseJSONBuffer", err)
		return nil, somethingError
	}
	operation := account.ServiceAccountInput{}

	accountID64, ok := jsonParsed.Path(`account_id`).Data().(float64)
	if !ok {
		return nil, fmt.Errorf("account_id is required param")
	}
	serviceID64, ok := jsonParsed.Path(`service_id`).Data().(float64)
	if !ok {
		return nil, fmt.Errorf("service_id is required param")
	}
	operationID64, ok := jsonParsed.Path(`operation_id`).Data().(float64) // id заказа
	if !ok {
		return nil, operationIdRequiredError
	}
	totalCost64, ok := jsonParsed.Path(`total_cost`).Data().(float64)
	if !ok {
		return nil, fmt.Errorf("total_cost is required param")
	}
	if totalCost64 <= 0 {
		return nil, fmt.Errorf("total_cost must be greater than zero")
	}

	operation.OperationID = int(operationID64)
	operation.AccountID = int(accountID64)
	operation.ServiceID = int(serviceID64)
	operation.TotalCost = float32(totalCost64)
	return &operation, nil
}
