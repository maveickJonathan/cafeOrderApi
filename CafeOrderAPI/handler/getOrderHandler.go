package handler

import (
	"net/http"

	"CafeOrderAPI/database_handler"
	"CafeOrderAPI/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type getOrderHandler struct {
	DB *gorm.DB
}

func NewGetOrderHandler(DB *gorm.DB) *fetchOrderHandler {
	return &fetchOrderHandler{DB}
}

func (h *fetchOrderHandler) GetOrder(c *gin.Context) {

	orderDB := database_handler.New(h.DB)

	order, errorDB := orderDB.GetAllOrder()

	if errorDB != nil {
		c.JSON(http.StatusBadRequest, errorDB)
		return
	}

	var responses []response.OrderResponse

	for i := 0; i < len(order); i++ {
		response := response.OrderResponse{
			Price:       order[i].Price,
			TotalPrice:  order[i].TotalPrice,
			ProductName: order[i].ProductName,
			Quantity:    order[i].Quantity,
		}
		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, responses)

}
