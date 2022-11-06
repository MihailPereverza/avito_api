package account

import (
	"avito_api/internal/db/model"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"log"
	"net/http"
)

func (ah *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	account, err := parseJSONCreateAccount(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = ah.accountService.CreateAccount(account)
	errMessage := operateCreateAccountError(account, err)
	if errMessage != "" {
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := []byte(fmt.Sprintf(
		`{"id":%d, "user_id":%d, "currency_id":%d, "balance":%f, "reserved_balance":%f}`,
		account.ID, account.User.ID, account.Currency.ID, account.Balance, account.ReservedBalance))

	if _, err := w.Write(resp); err != nil {
		log.Fatalf("Error happened in ResponseWriter Write. Err: %s", err)
	}
	return
}

func parseJSONCreateAccount(r *http.Request) (*model.Account, error) {
	account := model.Account{}
	jsonParsed, err := gabs.ParseJSONBuffer(r.Body)
	if err != nil {
		log.Printf("handle error: %s| in parseJSONAddBalance ParseJSONBuffer", err)
		return nil, somethingError
	}
	var ok bool

	userID64, ok := jsonParsed.Path("user_id").Data().(float64)
	if !ok {
		return nil, userIdRequiredError
	}
	account.User.ID = int(userID64)

	currencyID64, ok := jsonParsed.Path("currency_id").Data().(float64)
	if !ok {
		currencyID64 = -1
	}
	account.Currency.ID = int(currencyID64)
	return &account, nil
}

func operateCreateAccountError(account *model.Account, err error) string {
	errMessage := ""
	if err == nil {

	} else if err.Error() == duplicateError {
		errMessage = fmt.Sprintf("user already has %d", account.Currency.ID)
	} else if err.Error() == notFoundCurrency {
		errMessage = fmt.Sprintf("no found currency %d", account.Currency.ID)
	} else if err.Error() == notFoundUser {
		errMessage = fmt.Sprintf("not found user %d", account.User.ID)
	} else if err != nil {
		log.Printf("handle error: %s| on AccountHandler CreateAccount on call accountService", err)
		errMessage = `something error in request data`
	}
	return errMessage
}
