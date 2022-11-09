package account

type GetStatisticInput struct {
	AccountID       int
	OrderBy         string
	LastOperationID int
	LastPage        int
	NewPage         int
	OrderDirection  int
}
