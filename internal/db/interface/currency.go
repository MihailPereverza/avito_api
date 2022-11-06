package db_interface

type CurrencyDB interface {
	GetAllCurrency() (*map[string]int, error)
}
