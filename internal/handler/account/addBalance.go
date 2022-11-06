package account

import (
	"avito_api/internal/db/model"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"log"
	"net/http"
)

func (ah *AccountHandler) AddBalance(w http.ResponseWriter, r *http.Request) {
	account, amount, err := parseJSONAddBalance(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = ah.accountService.AddBalance(account, amount)
	errMessage := operateAddBalanceError(account, err)
	if errMessage != "" {
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := []byte(fmt.Sprintf(
		`{"id":%d, "user_id":%d, "currency_id":%d, "balance":%f}`,
		account.ID, account.User.ID, account.Currency.ID, account.Balance))
	if _, err := w.Write(resp); err != nil {
		log.Fatalf("Error happened in ResponseWriter Write. Err: %s", err)
	}
	return
}

func parseJSONAddBalance(r *http.Request) (*model.Account, float32, error) {
	account := model.Account{}
	jsonParsed, err := gabs.ParseJSONBuffer(r.Body)
	if err != nil {
		log.Printf("handle error: %s| in parseJSONAddBalance ParseJSONBuffer", err)
		return nil, -1, somethingError
	}
	var ok bool
	amount64, ok := jsonParsed.Path(`amount`).Data().(float64)
	if err != nil {
		return nil, -1, amountRequiredError
	}
	amount := float32(amount64)
	if amount <= 0 {
		return nil, -1, balanceGreaterZero
	}

	userID64, ok := jsonParsed.Path("user_id").Data().(float64)
	if !ok {
		return nil, -1, userIdRequiredError
	}
	account.User.ID = int(userID64)

	currencyID64, ok := jsonParsed.Path("currency_id").Data().(float64)
	if !ok {
		return nil, -1, currencyIDRequiredError
	}
	account.Currency.ID = int(currencyID64)
	return &account, amount, nil
}

func operateAddBalanceError(account *model.Account, err error) string {
	errMessage := ""
	if err == nil {

	} else if err.Error() == notFoundAccountToSet {
		errMessage = fmt.Sprintf(`not found account user_id %d, currency_id %d`, account.User.ID, account.Currency.ID)
	} else if err != nil {
		log.Printf("handle error: %s| on AccountHandler CreateAccount on call accountService", err)
		errMessage = `something error in request data`
	}
	return errMessage
}
