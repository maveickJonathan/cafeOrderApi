package handler

import (
	"fmt"
	"math"
	"net/http"
	"reflect"

	"CafeOrderAPI/database_handler"
	"CafeOrderAPI/request"
	"CafeOrderAPI/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type fetchOrderHandler struct {
	DB *gorm.DB
}

func NewFetchOrderHandler(DB *gorm.DB) *fetchOrderHandler {
	return &fetchOrderHandler{DB}
}

func (h *fetchOrderHandler) FetchOrder(c *gin.Context) {

	var request request.OrderRequest
	var response response.OrderResponse
	var quantity float64
	var errorVal string
	var priceflagError int
	errorMessages := []string{}

	err := c.ShouldBindJSON(&request)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

			if e.Field() == "Price" {
				priceflagError = 1
			}
		}
		//also Add the quantity error if it is error

		quantity, errorVal = ConvertAndValidateQuantity(request)
		if errorVal != "" {
			errorMessages = append(errorMessages, errorVal)
		}
	} else {
		quantity, errorVal = ConvertAndValidateQuantity(request)
		if errorVal != "" {
			errorMessages = append(errorMessages, errorVal)
		}
	}

	price, errorVal := ConvertAndValidatePrice(request, priceflagError)
	if errorVal != "" {
		errorMessages = append(errorMessages, errorVal)
	}

	response.Price = price
	response.Quantity = int(quantity)
	response.ProductName = request.ProductName
	response.TotalPrice = price * quantity

	if len(errorMessages) == 0 {
		orderDB := database_handler.New(h.DB)

		errorDB := orderDB.AddOrder(response)

		if errorDB != nil {
			errorMessages = append(errorMessages, errorDB.Error())
		}
	}

	if len(errorMessages) != 0 {
		c.JSON(http.StatusBadRequest, errorMessages)
		return
	}

	c.JSON(http.StatusOK, response)
}

func ConvertAndValidateQuantity(request request.OrderRequest) (float64, string) {
	quantity, err := getFloat(request.Quantity)
	if err != nil {
		errorMessage := "could not process quantity"
		return 0, errorMessage
	}

	//validate quantity is an integer
	valid := isFloatInt(quantity)

	if !valid {
		errorMessage := "Error on field Quantity, condition: quantity has to be integer"
		return 0, errorMessage
	}

	return quantity, ""
}

func ConvertAndValidatePrice(request request.OrderRequest, priceflagError int) (float64, string) {
	if priceflagError == 1 {
		return 0, ""
	}
	price, err := getFloat(request.Price)
	if err != nil {
		errorMessage := "could not process price"
		return 0, errorMessage
	}

	return price, ""
}

func getFloat(unk interface{}) (float64, error) {
	var floatType = reflect.TypeOf(float64(0))

	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}

func isFloatInt(floatValue float64) bool {
	return math.Mod(floatValue, 1.0) == 0
}
