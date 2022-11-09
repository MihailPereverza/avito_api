package account

import (
	"avito_api/internal/handler/inputModel/account"
	account2 "avito_api/internal/handler/outputModel/account"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"log"
	"net/http"
)

// @Summary Transfer
// @Tags account
// @Description transfer money from "from_account" to "to_account"
// @ID account-transfer
// @Accept json
// @Produce json
// @Param input body account.TransferInfo true "transfer info"
// @Success 200 {object} TransferOutput
// @failure 400 {object} outputModel.Error
// @Router /api/account/transfer [post]
func (ah *AccountHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	input, err := parseJSONTransfer(r)
	if err != nil {
		jsonError(w, err)
		return
	}

	transferOutput, err := ah.accountService.Transfer(input)
	errMessage := operateTransferError(input, err)
	if errMessage != "" {
		jsonError(w, fmt.Errorf(errMessage))
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, err := respJSONTransfer(transferOutput)
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

func parseJSONTransfer(r *http.Request) (*account.TransferInfo, error) {
	jsonParsed, err := gabs.ParseJSONBuffer(r.Body)
	if err != nil {
		log.Printf("handle error: %s| in parseJSONAddBalance ParseJSONBuffer", err)
		return nil, somethingError
	}

	transferInfo := account.TransferInfo{}
	fromAccountID64, ok := jsonParsed.Path(`from_account_id`).Data().(float64)
	if !ok {
		return nil, fmt.Errorf("from_account_id is required param")
	}
	toAccountID64, ok := jsonParsed.Path(`to_account_id`).Data().(float64)
	if !ok {
		return nil, fmt.Errorf("to_account_id is required param")
	}
	amount64, ok := jsonParsed.Path(`amount`).Data().(float64)
	if !ok {
		return nil, fmt.Errorf("amount is required param")
	}
	if amount64 <= 0 {
		return nil, fmt.Errorf("total_cost must be greater than zero")
	}

	transferInfo.FromAccountID = int(fromAccountID64)
	transferInfo.ToAccountID = int(toAccountID64)
	transferInfo.Amount = float32(amount64)
	return &transferInfo, nil
}

func respJSONTransfer(transferOutput *account2.TransferOutput) (*gabs.Container, error) {
	resp := gabs.New()
	setFieldError := fmt.Errorf("cant generate response json")
	_, err := resp.Set(transferOutput.FromAccount.ID, "from_account", "account_id")
	if err != nil {
		log.Printf("Error happened in set json. Err: %s", err)
		return nil, setFieldError
	}
	_, err = resp.Set(transferOutput.FromAccount.Balance, "from_account", "balance")
	if err != nil {
		return nil, setFieldError
	}
	_, err = resp.Set(transferOutput.FromAccount.ReservedBalance, "from_account", "reserved_balance")
	if err != nil {
		return nil, setFieldError
	}

	_, err = resp.Set(transferOutput.ToAccount.ID, "to_account", "account_id")
	if err != nil {
		log.Printf("Error happened in set json. Err: %s", err)
		return nil, setFieldError
	}
	_, err = resp.Set(transferOutput.ToAccount.Balance, "to_account", "balance")
	if err != nil {
		return nil, setFieldError
	}
	_, err = resp.Set(transferOutput.ToAccount.ReservedBalance, "to_account", "reserved_balance")
	if err != nil {
		return nil, setFieldError
	}

	return resp, nil
}

func operateTransferError(account *account.TransferInfo, err error) string {
	errMessage := ""
	if err == nil {

	} else if err.Error() == notFoundAccountToSet {
		errMessage = fmt.Sprintf(`not found couple  (from_account_id, to_account_id) (%d, %d)`, account.FromAccountID, account.ToAccountID)
	} else if err.Error() == enoughMoney {
		errMessage = fmt.Sprintf(`enough money`)
	} else if err != nil {
		log.Printf("handle error: %s| on AccountHandler reserveAccount on call accountService", err)
		errMessage = `something error in request data`
	}
	return errMessage
}
