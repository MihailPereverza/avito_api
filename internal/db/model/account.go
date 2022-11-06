package model

type Account struct {
	ID              int      `json:"id"`
	User            User     `json:"user"`
	Currency        Currency `json:"currency"`
	Balance         float32  `json:"balance"`
	ReservedBalance float32  `json:"reserved_balance"`
}
