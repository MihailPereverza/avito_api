package service

import (
	"avito_api/internal/db/interface"
	"avito_api/internal/db/model"
	account2 "avito_api/internal/handler/inputModel/account"
	"avito_api/internal/handler/outputModel/account"
	"avito_api/internal/service/interface"
	"fmt"
)

type AccountService struct {
	accountDB   db_interface.AccountDB
	operationDB db_interface.OperationDB
}

func NewAccountService(
	dbAccount db_interface.AccountDB,
	dbOperation db_interface.OperationDB,
) service_interface.AccountServiceInterface {
	return &AccountService{
		dbAccount,
		dbOperation,
	}
}

func (us *AccountService) CreateAccount(account *model.DBCreateAccount) error {
	return us.accountDB.CreateAccount(account)
}

func (us *AccountService) Transfer(info *account2.TransferInfo) (*account.TransferOutput, error) {
	fromAccount, toAccount, err := us.accountDB.Transfer(&model.DBTransferInfo{TransferInfo: *info})
	return &account.TransferOutput{FromAccount: *fromAccount, ToAccount: *toAccount}, err
}

func (us *AccountService) AddBalance(accountInput *account2.AddBalanceInput) (*account.AddBalanceOutput, error) {
	exists, err := us.accountDB.IsExists(accountInput.ID)
	if err != nil {
		return nil, err
	} else if !exists {
		err = us.CreateAccount(&model.DBCreateAccount{ID: accountInput.ID})
	}
	if err != nil {
		return nil, err
	}
	balanceAccount, err := us.accountDB.AddBalance(&model.DBAddBalanceInput{ID: accountInput.ID, Amount: accountInput.Amount})
	if err != nil {
		return nil, err
	}
	return &account.AddBalanceOutput{ID: balanceAccount.ID, Balance: balanceAccount.Balance}, nil
}

func (us *AccountService) ReserveBalance(operation *account2.ServiceAccountInput) (*model.Account, error) {
	dbInput := model.DBReserveBalance{ServiceAccountInput: *operation}
	return us.accountDB.ReserveBalance(&dbInput)
}

func (us *AccountService) UnReserveBalance(operation *account2.ServiceAccountInput) (*model.Account, error) {
	inputDB := model.DBBaseOperationInfo{ServiceAccountInput: *operation}
	return us.accountDB.UnReserveBalance(&inputDB)
}

func (us *AccountService) ApproveDebiting(operation *account2.ServiceAccountInput) (*account.ApproveDebitingOutput, error) {
	dbApprove := model.DBBaseOperationInfo{ServiceAccountInput: *operation}
	companyBalance, err := us.operationDB.ApproveDebiting(&dbApprove)
	if err != nil {
		return nil, err
	}
	return &account.ApproveDebitingOutput{CompanyBalance: companyBalance}, nil
}

func (us *AccountService) GetBalance(accountID int) (*model.Account, error) {
	return us.accountDB.GetBalance(accountID)
}

func (us *AccountService) GetStatistic(input *account2.GetStatisticInput) ([]account.GetStatisticOutput, error) {
	direction := 0
	count := 1
	if input.NewPage > input.LastPage {
		direction = 1
		count *= input.NewPage - input.LastPage
	} else if input.NewPage == input.LastPage {
		return nil, fmt.Errorf("newPage and lastPage must be different")
	} else if input.NewPage < input.LastPage {
		direction = -1
		count *= input.LastPage - input.NewPage
	}
	dbInput := model.DBGetStatistic{
		AccountID:       input.AccountID,
		OrderBy:         input.OrderBy,
		OrderDirection:  input.OrderDirection,
		Direction:       direction,
		LastOperationID: input.LastOperationID,
		Count:           count,
	}
	var output []account.GetStatisticOutput
	dbOperations, err := us.accountDB.GetStatistic(&dbInput)
	if err != nil {
		return nil, err
	}
	for _, v := range dbOperations {
		output = append(output, account.GetStatisticOutput{DBGetStatisticOutput: v})
	}

	return output, nil
}
