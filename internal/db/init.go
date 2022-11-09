package db

import (
	"database/sql"
	"fmt"
)

func Init(db *sql.DB) error {
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

func createAccountTable(db *sql.DB) error {
	account, err := db.Prepare(`CREATE TABLE IF NOT EXISTS account(
    	account_id SERIAL PRIMARY KEY,
    	balance Decimal(13,4) NOT NULL DEFAULT 0 CHECK (balance >= 0),
    	reserved_balance Decimal(13,4) NOT NULL DEFAULT 0 CHECK (reserved_balance >= 0)
	)`)
	if err != nil {
		return err
	}
	if _, err = account.Exec(); err != nil {
		return err
	}

	// если есть ошибка -> "аккаунт компании" уже создан
	_, err = db.Exec(`INSERT INTO account(account_id) VALUES (0);`)
	if err != nil {
		fmt.Printf("try add company in table error: %s", err)
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
    	operation_id INTEGER PRIMARY KEY,
    	account_id INTEGER REFERENCES account(account_id) NOT NULL,
    	status_id INTEGER REFERENCES operation_status(status_id) NOT NULL,
    	service_id INTEGER REFERENCES service(service_id) NOT NULL,
    	count INTEGER NOT NULL DEFAULT 1,
    	total_cost Decimal(13,4) NOT NULL,
    	create_time TIMESTAMP DEFAULT now()::timestamp
	)`)
	if err != nil {
		return err
	}
	if _, err = operation.Exec(); err != nil {
		return err
	}
	return nil
}
