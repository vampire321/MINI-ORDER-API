// http concerns
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mini-order-api/internal/model"
	"mini-order-api/internal/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {  //Provides a clean way to create a new OrderHandler with its required OrderService.Keeps your code modular and testable.
	return &OrderHandler{service: service}
}

func (h *OrderHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/orders", h.Create)
	r.GET("/orders", h.GetAll)
}

func (h *OrderHandler) Create(c *gin.Context) {

	var order model.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateOrder(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetAll(c *gin.Context) {

	orders, err := h.service.GetOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}