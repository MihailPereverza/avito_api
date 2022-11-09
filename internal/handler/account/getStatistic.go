package account

import (
	"avito_api/internal/handler/inputModel/account"
	account2 "avito_api/internal/handler/outputModel/account"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// @Summary GetBalance
// @Tags account
// @Description get transaction statisctic
// @ID account-get-statistic
// @Accept json
// @Produce json
// @Param account_id path int true "account_id"
// @Success 200 {object} []account.GetStatisticOutput
// @failure 400 {object} outputModel.Error
// @Router /api/account/{account_id}/statistic [get]
func (ah *AccountHandler) GetStatistic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	input, err := parseParamsStatistic(r)
	if err != nil {
		jsonError(w, err)
		return
	}

	operations, err := ah.accountService.GetStatistic(input)
	errMessage := operateGetStatisticError(input.AccountID, err)
	if errMessage != "" {
		jsonError(w, fmt.Errorf(errMessage))
		return
	}
	if len(operations) == 0 {
		// другие параметры либо обрабатываются (order, direction),
		// либо не могут выдать нулевой результат: (lpage, npage) - влияют на count
		jsonError(w, fmt.Errorf("incorrect one of params: operations"))
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := generateJSONStatistic(operations)
	_, err = resp.Set(operations[len(operations)-1].OperationID, "last_operation_id")
	if resp == nil {
		log.Printf("Error happened in set json. Err: %s", err)
		jsonError(w, fmt.Errorf("cant generate response json"))
		return
	}
	if _, err = w.Write(resp.Bytes()); err != nil {
		log.Printf("Error happened in ResponseWriter Write. Err: %s", err)
		jsonError(w, fmt.Errorf("cant send response json"))
		return
	}
	return
}

func parseParamsStatistic(r *http.Request) (*account.GetStatisticInput, error) {
	input := account.GetStatisticInput{}
	var err error

	input.AccountID, err = strconv.Atoi(mux.Vars(r)["accountID"])
	if err != nil {
		return nil, fmt.Errorf("cant find account_id")
	}
	input.OrderBy = r.URL.Query().Get("order")
	if input.OrderBy != "cost" && input.OrderBy != "date" {
		return nil, fmt.Errorf("cant order by %s", input.OrderBy)
	}

	// lastOperationId: -1 - начальная страница, 0-(+inf) - позиция
	//для начала отсчета пагинации
	input.LastOperationID = -2
	lastOperationIdStr := r.URL.Query().Get("operation")
	if lastOperationIdStr == "" {
		input.LastOperationID = -1
	} else {
		input.LastOperationID, err = strconv.Atoi(lastOperationIdStr)
	}
	if err != nil || input.LastOperationID < -1 {
		fmt.Printf("AccountHandler.GetStatistic.strconv.Atoi handle error: %s", err)
		return nil, fmt.Errorf("cant operate operation_id")
	}

	input.LastPage, err = strconv.Atoi(r.URL.Query().Get("lpage"))
	if err != nil || input.LastPage < 0 {
		return nil, fmt.Errorf("cant operate lpage")
	}
	input.NewPage, err = strconv.Atoi(r.URL.Query().Get("npage"))
	if err != nil || input.NewPage < 0 {
		return nil, fmt.Errorf("cant operate npage")
	}
	input.OrderDirection, err = strconv.Atoi(r.URL.Query().Get("direction"))
	if err != nil || (input.OrderDirection != 1 && input.OrderDirection != 0) {
		return nil, fmt.Errorf("direction has been 0 (DESK) or 1 (ASK), but got %d", input.OrderDirection)
	}

	return &input, nil
}

func generateJSONStatistic(operations []account2.GetStatisticOutput) *gabs.Container {
	resp := gabs.New()
	_, err := resp.Array("result")
	if err != nil {
		log.Printf("error generateJSONStatistic.resp.Array: %s", err)
		return nil
	}
	for _, operation := range operations {
		jsonOperation := gabs.New()
		_, err = jsonOperation.Set(operation.OperationID, "order_id")
		if err != nil {
			return nil
		}
		_, err = jsonOperation.Set(operation.StatusTitle, "status")
		if err != nil {
			return nil
		}
		_, err = jsonOperation.Set(operation.ServiceTitle, "service_title")
		if err != nil {
			return nil
		}
		_, err = jsonOperation.Set(operation.ServiceDescription, "service_description")
		if err != nil {
			return nil
		}
		_, err = jsonOperation.Set(operation.TotalCost, "total_cost")
		if err != nil {
			return nil
		}
		_, err = jsonOperation.Set(operation.CreateTime, "create_time")
		if err != nil {
			return nil
		}
		err = resp.ArrayAppend(jsonOperation, "result")
		if err != nil {
			return nil
		}
	}
	return resp
}

func operateGetStatisticError(accountID int, err error) string {
	errMessage := ""
	if err == nil {

	} else if err.Error() == notFoundAccountToSet {
		errMessage = fmt.Sprintf(`not found account %d`, accountID)
	} else if err.Error() == notFoundAccountOperation {
		errMessage = `account has not operations`
	} else if err != nil {
		log.Printf("handle error: %s| on AccountHandler CreateAccount on call accountService", err)
		errMessage = `something error in request data`
	}
	return errMessage
}
