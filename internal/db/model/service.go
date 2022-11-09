package model

type Service struct {
	ID          int     `json:"id"`
	OrderID     int     `json:"order_id"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}
