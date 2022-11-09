package account

type ServiceAccountInput struct {
	AccountID   int     `json:"account_id"`
	ServiceID   int     `json:"service_id"`
	OperationID int     `json:"operation_id"`
	TotalCost   float32 `json:"total_cost"`
}
