package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"mini-order-api/internal/config"
	"mini-order-api/internal/database"
	"mini-order-api/internal/handler"
	"mini-order-api/internal/repository"
	"mini-order-api/internal/service"
)

func main() {

	cfg := config.LoadConfig()

	db,err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewOrderRepository(db)

	service := service.NewOrderService(repo)

	handler := handler.NewOrderHandler(service)

	r := gin.Default()

	handler.RegisterRoutes(r)

	r.Run(":8080")
}