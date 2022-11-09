package account

import "avito_api/internal/db/model"

type TransferOutput struct {
	FromAccount model.Account `json:"from_account"`
	ToAccount   model.Account `json:"to_account"`
}
