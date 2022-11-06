package currency

import (
	"avito_api/internal/db/interface"
	"database/sql"
)

var allCurrency = map[string]int{}

type DBCurrency struct {
	db *sql.DB
}

func NewDBCurrency(db *sql.DB) db_interface.CurrencyDB {
	return &DBCurrency{db}
}

func (D *DBCurrency) GetAllCurrency() (*map[string]int, error) {
	if len(allCurrency) != 0 {
		return &allCurrency, nil
	}
	rows, err := D.db.Query(`SELECT currency_id, symbol FROM currency`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var curID int
		var curSymb string
		if err = rows.Scan(&curID, &curSymb); err != nil {
			return nil, err
		}
		allCurrency[curSymb] = curID
	}
	return &allCurrency, err
}
