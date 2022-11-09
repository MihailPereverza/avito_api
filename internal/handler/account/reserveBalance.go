package account

import (
	"avito_api/internal/db/model"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"log"
	"net/http"
)

// @Summary ReserveBalance
// @Tags account
// @Description reserve balance from account
// @ID account-reserve-balance
// @Accept json
// @Produce json
// @Param input body account.ServiceAccountInput true "account info"
// @Success 200 {object} model.Account
// @failure 400 {object} outputModel.Error
// @Router /api/account/balance/reserve [post]
func (ah *AccountHandler) ReserveBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	input, err := parseJSONServiceAccount(r)
	if err != nil {
		jsonError(w, err)
		return
	}

	account, err := ah.accountService.ReserveBalance(input)
	errMessage := operateReserveBalanceError(account, err)
	if errMessage != "" {
		jsonError(w, fmt.Errorf(errMessage))
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, err := respJSONAccount(account)
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

func respJSONAccount(account *model.Account) (*gabs.Container, error) {
	resp := gabs.New()
	setFieldError := fmt.Errorf("cant generate response json")
	_, err := resp.Set(account.ID, "account_id")
	if err != nil {
		log.Printf("Error happened in set json. Err: %s", err)
		return nil, setFieldError
	}
	_, err = resp.Set(account.Balance, "balance")
	if err != nil {
		return nil, setFieldError
	}
	_, err = resp.Set(account.ReservedBalance, "reserved_balance")
	if err != nil {
		return nil, setFieldError
	}
	return resp, nil
}

func operateReserveBalanceError(account *model.Account, err error) string {
	errMessage := ""
	if err == nil {

	} else if err.Error() == notFoundAccountToUpdate {
		errMessage = fmt.Sprintf(`not found account_id %d`, account.ID)
	} else if err.Error() == insufficientFunds {
		errMessage = fmt.Sprintf(`insufficient funds`)
	} else if err.Error() == notFoundOperation {
		errMessage = fmt.Sprintf(`not found service`)
	} else if err.Error() == operationAlreadyExists {
		errMessage = fmt.Sprintf(`order_id=%d already exists`, account.ID)
	} else if err != nil {
		log.Printf("handle error: %s| on AccountHandler reserveAccount on call accountService", err)
		errMessage = `something error in request data`
	}
	return errMessage
}
