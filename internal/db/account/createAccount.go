package account

import (
	"avito_api/internal/db/model"
	"fmt"
)

func (D *DBAccount) CreateAccount(account *model.DBCreateAccount) error {
	if account.ID < 1 {
		return fmt.Errorf("cant create account with ID < 1")
	}
	query := `INSERT INTO account(account_id) VALUES($1)`
	_, err := D.db.Exec(query, account.ID)
	return err
}
