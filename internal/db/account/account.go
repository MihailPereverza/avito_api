package account

import (
	"avito_api/internal/db/interface"
	"avito_api/internal/db/model"
	"database/sql"
)

type DBAccount struct {
	db *sql.DB
}

func NewDBAccount(db *sql.DB) db_interface.AccountDB {
	return &DBAccount{db}
}

func (D *DBAccount) CreateAccount(account *model.Account) (*model.Account, error) {
	query := `INSERT INTO account(user_id, currency_id) VALUES ($1, $2) RETURNING account_id`
	err := D.db.QueryRow(query, account.User.ID, account.Currency.ID).Scan(&account.ID)
	return account, err
}

func (D *DBAccount) AddBalance(account *model.Account, amount float32) (*model.Account, error) {
	query := `UPDATE account
    	SET balance = balance + $1
		WHERE user_id = $2 AND currency_id = $3 
		RETURNING account_id, balance`
	err := D.db.QueryRow(
		query, amount, account.User.ID, account.Currency.ID,
	).Scan(&account.ID, &account.Balance)
	return account, err
}
