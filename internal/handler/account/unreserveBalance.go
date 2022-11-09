package account

import (
	"avito_api/internal/db/model"
	"avito_api/internal/handler/inputModel/account"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"log"
	"net/http"
)

// @Summary UnReserveBalance
// @Tags account
// @Description unreserve balance from account
// @ID account-unreserve-balance
// @Accept json
// @Produce json
// @Param input body account.ServiceAccountInput true "operation info"
// @Success 200 {object} model.Account
// @failure 400 {object} outputModel.Error
// @Router /api/account/balance/unreserve [post]
func (ah *AccountHandler) UnReserveBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	operation, err := parseJSONServiceAccount(r)
	if err != nil {
		jsonError(w, err)
		return
	}

	accountInfo, err := ah.accountService.UnReserveBalance(operation)
	errMessage := operateUnReserveBalanceError(operation, err)
	if errMessage != "" {
		jsonError(w, fmt.Errorf(errMessage))
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, err := jsonResp(accountInfo)
	if err != nil {
		jsonError(w, err)
		return
	}

	if _, err = w.Write(resp.Bytes()); err != nil {
		log.Printf("Error happened in ResponseWriter Write. Err: %s", err)
		jsonError(w, fmt.Errorf("cant send response json"))
		return
	}
	return
}

func jsonResp(account *model.Account) (*gabs.Container, error) {
	resp := gabs.New()
	jsonSetError := fmt.Errorf("cant generate response json")
	_, err := resp.Set(account.ID, "account_id")
	if err != nil {
		return nil, jsonSetError
	}
	_, err = resp.Set(account.Balance, "balance")
	if err != nil {
		return nil, jsonSetError
	}
	_, err = resp.Set(account.ReservedBalance, "reserved_balance")
	if err != nil {
		return nil, jsonSetError
	}
	return resp, err
}

func operateUnReserveBalanceError(operation *account.ServiceAccountInput, err error) string {
	errMessage := ""
	if err == nil {

	} else if err.Error() == notFoundAccountToSet {
		errMessage = fmt.Sprintf(`not found order_id %d`, operation.OperationID)
	} else if err.Error() == operationCantUnReserve {
		errMessage = fmt.Sprintf(`order_id=%d cant unreserve`, operation.OperationID)
	} else if err.Error() == operationAlreadyExists {
		errMessage = fmt.Sprintf(`order_id=%d already exists`, operation.OperationID)
	} else if err != nil {
		log.Printf("handle error: %s| on AccountHandler reserveAccount on call accountService", err)
		errMessage = `something error in request data`
	}
	return errMessage
}
