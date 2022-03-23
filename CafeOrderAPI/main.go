package main

import (
	"CafeOrderAPI/handler"
	"CafeOrderAPI/model"
	"log"

	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dbURL := "host=database user=pg password=pass dbname=cafe_order port=5432"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.Order{})

	fmt.Println("Database connection succeed")

	router := gin.Default()

	fetchOrderHandler := handler.NewFetchOrderHandler(db)
	getOrderHandler := handler.NewGetOrderHandler(db)

	router.POST("/createOrder", fetchOrderHandler.FetchOrder)
	router.GET("/getOrder", getOrderHandler.GetOrder)

	router.Run(":8080")
}
