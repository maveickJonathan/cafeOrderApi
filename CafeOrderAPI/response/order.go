package response

type OrderResponse struct {
	Price       float64 `json:"price"`
	TotalPrice  float64 `json:"total_price"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
}
