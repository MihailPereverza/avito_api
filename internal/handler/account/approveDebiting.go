package account

import (
	"avito_api/internal/handler/inputModel/account"
	"errors"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"log"
	"net/http"
)

// @Summary ApproveDebiting
// @Tags account
// @Description approve-debiting
// @ID account-approve-debiting
// @Accept json
// @Produce json
// @Param input body account.ServiceAccountInput true "account_id and balance"
// @Success 200 {object} ApproveDebitingOutput
// @failure 400 {object} outputModel.Error
// @Router /api/operation/approve [post]
func (ah *AccountHandler) ApproveDebiting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	operation, err := parseJSONServiceAccount(r)
	if err != nil {
		jsonError(w, err)
		return
	}
	companyBalance, err := ah.accountService.ApproveDebiting(operation)
	errMessage := operateApproveDebitingError(operation, err)
	if errMessage != "" {
		jsonError(w, errors.New(errMessage))
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := gabs.New()
	_, err = resp.Set(companyBalance.CompanyBalance, "company_balance")
	if err != nil {
		log.Fatalf("Error happened in set json. Err: %s", err)
	}

	if _, err := w.Write(resp.Bytes()); err != nil {
		log.Fatalf("Error happened in ResponseWriter Write. Err: %s", err)
	}
	return
}

func operateApproveDebitingError(operation *account.ServiceAccountInput, err error) string {
	errMessage := ""
	if err == nil {

	} else if err.Error() == notFoundAccountToSet {
		errMessage = fmt.Sprintf(`not found operation_id %d`, operation.OperationID)
	} else if err.Error() == notAffectedRows {
		errMessage = fmt.Sprintf(`not found operation_id %d`, operation.OperationID)
	} else if err.Error() == operationAlreadyApproved {
		errMessage = fmt.Sprintf(`operation_id %d already operated`, operation.OperationID)
	} else if err != nil {
		log.Printf("handle error: %s| on AccountHandler reserveAccount on call accountService", err)
		errMessage = `something error in request data`
	}
	return errMessage
}
