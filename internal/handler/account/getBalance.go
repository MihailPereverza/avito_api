package account

import (
	"avito_api/internal/db/model"
	"errors"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"log"
	"net/http"
)

// @Summary GetBalance
// @Tags account
// @Description get balance
// @ID account-get-balance
// @Accept json
// @Produce json
// @Param input body string true "account_id" SchemaExample({"account_id": 1234})
// @Success 200 {object} ApproveDebitingOutput
// @failure 400 {object} outputModel.Error
// @Router /api/account/balance [post]
func (ah *AccountHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	accountID, err := parseJSONAddGetBalance(r)
	if err != nil {
		jsonError(w, err)
		return
	}

	account, err := ah.accountService.GetBalance(accountID)
	errMessage := operateGetBalanceError(accountID, err)
	if errMessage != "" {
		jsonError(w, errors.New(errMessage))
		return
	}

	resp := respJSONGetBalance(account)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp.Bytes()); err != nil {
		log.Fatalf("Error happened in ResponseWriter Write. Err: %s", err)
	}
	return
}

func respJSONGetBalance(account *model.Account) *gabs.Container {
	resp := gabs.New()
	_, err := resp.Set(account.ID, "account_id")
	if err != nil {
		log.Fatalf("Error happened in set json. Err: %s", err)
	}
	_, err = resp.Set(account.Balance, "balance")
	if err != nil {
		log.Fatalf("Error happened in set json. Err: %s", err)
	}
	_, err = resp.Set(account.ReservedBalance, "reserved_balance")
	if err != nil {
		log.Fatalf("Error happened in set json. Err: %s", err)
	}
	return resp
}

func parseJSONAddGetBalance(r *http.Request) (int, error) {
	accountID := -1
	jsonParsed, err := gabs.ParseJSONBuffer(r.Body)
	if err != nil {
		return accountID, fmt.Errorf("cant parse json")
	}
	var ok bool
	accountID64, ok := jsonParsed.Path(`account_id`).Data().(float64)
	if !ok {
		return accountID, operationIdRequiredError
	}
	accountID = int(accountID64)

	return accountID, nil
}

func operateGetBalanceError(accountID int, err error) string {
	errMessage := ""
	if err == nil {

	} else if err.Error() == notFoundAccountToSet {
		errMessage = fmt.Sprintf(`not found account %d`, accountID)
	} else if err != nil {
		log.Printf("handle error: %s| on AccountHandler CreateAccount on call accountService", err)
		errMessage = `something error in request data`
	}
	return errMessage
}
