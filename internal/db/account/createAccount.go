package account

import "avito_api/internal/db/model"

func (D *DBAccount) CreateAccount(account *model.DBCreateAccount) error {
	query := `INSERT INTO account(account_id) VALUES($1)`
	_, err := D.db.Exec(query, account.ID)
	return err
}
