package model

import "time"

type Order struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Price       float64 `json:"price"`
	TotalPrice  float64 `json:"total_price"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
