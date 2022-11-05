package db

import (
	"database/sql"
)

func Init(db *sql.DB) error {
	if err := createUserTable(db); err != nil {
		return err
	}
	if err := createCurrencyTable(db); err != nil {
		return err
	}
	if err := createAccountTable(db); err != nil {
		return err
	}
	if err := createServiceTable(db); err != nil {
		return err
	}
	if err := createOperationStatusTable(db); err != nil {
		return err
	}
	if err := createAccountOperationTable(db); err != nil {
		return err
	}
	return nil
}

func createUserTable(db *sql.DB) error {
	// создаем последовательность для айдишников пользователей
	userSeq, err := db.Prepare(`CREATE SEQUENCE IF NOT EXISTS user_id_seq START 1373718781;`)
	if err != nil {
		return err
	}
	if _, err = userSeq.Exec(); err != nil {
		return err
	}

	// создаем таблицу пользователей
	users, err := db.Prepare(`CREATE TABLE  IF NOT EXISTS users(
		user_id INTEGER PRIMARY KEY DEFAULT nextval('user_id_seq'),
		name VARCHAR(64) NOT NULL,
		email VARCHAR(64) NOT NULL);`)
	if err != nil {
		return err
	}
	if _, err = users.Exec(); err != nil {
		return err
	}

	return nil
}

func createCurrencyTable(db *sql.DB) error {
	currency, err := db.Prepare(`CREATE TABLE IF NOT EXISTS currency(
    	currency_id SERIAL PRIMARY KEY,
    	name VARCHAR(24) NOT NULL,
    	symbol VARCHAR(8) NOT NULL,
		UNIQUE(name),
    	UNIQUE(symbol)
	)`)
	if err != nil {
		return err
	}
	if _, err = currency.Exec(); err != nil {
		return err
	}
	return nil
}

func createAccountTable(db *sql.DB) error {
	account, err := db.Prepare(`CREATE TABLE IF NOT EXISTS account(
    	account_id SERIAL PRIMARY KEY,
    	user_id INTEGER REFERENCES users(user_id) NOT NULL,
    	currency_id INTEGER REFERENCES currency(currency_id) NOT NULL,
    	balance Decimal(13,4) NOT NULL DEFAULT 0,
    	reserved_balance Decimal(13,4) NOT NULL DEFAULT 0,
    	UNIQUE (user_id, currency_id)
	)`)
	if err != nil {
		return err
	}
	if _, err = account.Exec(); err != nil {
		return err
	}
	return nil
}

func createServiceTable(db *sql.DB) error {
	service, err := db.Prepare(`CREATE TABLE IF NOT EXISTS service(
    	service_id SERIAL PRIMARY KEY,
    	title VARCHAR(128) NOT NULL,
    	price DECIMAL(13, 4) NOT NULL CHECK (price > 0),
    	description TEXT NOT NULL,
    	UNIQUE (title)
	)`)
	if err != nil {
		return err
	}
	if _, err = service.Exec(); err != nil {
		return err
	}
	return nil
}

func createOperationStatusTable(db *sql.DB) error {
	statuses, err := db.Prepare(`CREATE TABLE IF NOT EXISTS operation_status(
    	status_id SERIAL PRIMARY KEY,
    	title VARCHAR(32) NOT NULL
	)`)
	if err != nil {
		return err
	}
	if _, err = statuses.Exec(); err != nil {
		return err
	}
	return nil
}

func createAccountOperationTable(db *sql.DB) error {
	operation, err := db.Prepare(`CREATE TABLE IF NOT EXISTS account_operation(
    	operation_id SERIAL PRIMARY KEY,
    	account_id INTEGER REFERENCES account(account_id) NOT NULL,
    	status_id INTEGER REFERENCES operation_status(status_id) NOT NULL,
    	user_id INTEGER REFERENCES users(user_id) NOT NULL,
    	service_id INTEGER REFERENCES service(service_id) NOT NULL,
    	count INTEGER NOT NULL DEFAULT 1,
    	total_cost Decimal(13,4) NOT NULL
	)`)
	if err != nil {
		return err
	}
	if _, err = operation.Exec(); err != nil {
		return err
	}
	return nil
}
