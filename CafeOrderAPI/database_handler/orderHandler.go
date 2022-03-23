package database_handler

import (
	"CafeOrderAPI/model"
	"CafeOrderAPI/response"
	"fmt"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

func (h handler) GetAllOrder() ([]model.Order, error) {
	var orders []model.Order

	if result := h.DB.Find(&orders); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return orders, nil
}

func (h handler) AddOrder(response response.OrderResponse) error {
	fmt.Println("AddOrder")
	order := model.Order{
		Price:       response.Price,
		TotalPrice:  response.TotalPrice,
		ProductName: response.ProductName,
		Quantity:    response.Quantity,
	}
	fmt.Println(order)
	fmt.Println("masukkin data ke db")
	// Append to the Books table
	if result := h.DB.Create(&order); result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}
