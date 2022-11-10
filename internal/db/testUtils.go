package db

import (
	"avito_api/internal/config"
	"database/sql"
	"fmt"
)

func InitTestDB() (*sql.DB, error) {
	db, err := NewPostgres(config.GetDBConfig())
	if err != nil {
		return nil, fmt.Errorf("cant connect to test_db %w", err)
	}
	db.Exec(`DROP SCHEMA public CASCADE;`)
	_, err = db.Exec(`CREATE SCHEMA public;`)
	if err != nil {
		return nil, err
	}
	err = Init(db)
	if err != nil {
		return nil, fmt.Errorf("cant init test_db %w", err)
	}
	return db, nil
}
