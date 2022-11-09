package account

import (
	"avito_api/internal/handler/inputModel/account"
	account2 "avito_api/internal/handler/outputModel/account"
	"errors"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"log"
	"net/http"
)

// @Summary AddBalance
// @Tags account
// @Description add balance on account
// @ID account-add-balance
// @Accept json
// @Produce json
// @Param input body account.AddBalanceInput true "account_id and balance"
// @Success 200 {object} AddBalanceOutput
// @failure 400 {object} outputModel.Error
// @Router /api/account/balance/add [post]
func (ah *AccountHandler) AddBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	accountInput, err := parseJSONAddBalance(r)
	if err != nil {
		jsonError(w, err)
		return
	}
	accountOutput, err := ah.accountService.AddBalance(accountInput)
	errMessage := operateAddBalanceError(accountOutput, err)
	if errMessage != "" {
		jsonError(w, errors.New(errMessage))
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := respJSONAddBalance(accountOutput)
	if _, err = w.Write(resp.Bytes()); err != nil {
		log.Fatalf("Error happened in ResponseWriter Write. Err: %s", err)
	}
	return
}

func respJSONAddBalance(account *account2.AddBalanceOutput) *gabs.Container {
	resp := gabs.New()
	_, err := resp.Set(account.ID, "account_id")
	if err != nil {
		log.Fatalf("Error happened in set json. Err: %s", err)
	}
	_, err = resp.Set(account.Balance, "balance")
	if err != nil {
		log.Fatalf("Error happened in set json. Err: %s", err)
	}

	return resp
}

func parseJSONAddBalance(r *http.Request) (*account.AddBalanceInput, error) {
	amountAccount := account.AddBalanceInput{}
	jsonParsed, err := gabs.ParseJSONBuffer(r.Body)
	if err != nil {
		log.Printf("handle error: %s| in parseJSONAddBalance ParseJSONBuffer", err)
		return nil, somethingError
	}
	var ok bool
	amount64, ok := jsonParsed.Path(`amount`).Data().(float64)
	if amount64 == 0 {
		return nil, amountRequiredError
	}
	amount := float32(amount64)
	if amount < 0 {
		return nil, balanceGreaterZero
	}
	amountAccount.Amount = amount

	userID64, ok := jsonParsed.Path("account_id").Data().(float64)
	if !ok {
		return nil, accountIdRequiredError
	}
	amountAccount.ID = int(userID64)

	return &amountAccount, nil
}

func operateAddBalanceError(account *account2.AddBalanceOutput, err error) string {
	errMessage := ""
	if err == nil {

	} else if err.Error() == notFoundAccountToSet {
		errMessage = fmt.Sprintf(`not found account %d`, account.ID)
	} else if err != nil {
		log.Printf("handle error: %s| on AccountHandler CreateAccount on call accountService", err)
		errMessage = `something error in request data`
	}
	return errMessage
}
