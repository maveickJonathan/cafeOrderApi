package request

type OrderRequest struct {
	Price       interface{} `json:"price" binding:"required,number"`
	ProductName string      `json:"product_name" binding:"required"`
	Quantity    interface{} `json:"quantity" binding:"required,number"`
}
