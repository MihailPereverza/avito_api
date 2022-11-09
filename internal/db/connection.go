package db

import (
	"avito_api/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
)

func NewPostgres(cfg *config.DBConfig) (*sql.DB, error) {
	dbURI := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)
	db, err := sql.Open(cfg.Driver, dbURI)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("couldn't ping: postgres", err)
		return nil, err
	}
	if err = Init(db); err != nil {
		log.Println("couldn't init db (create tables)", err)
		return nil, err
	}

	return db, nil
}
