package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/services"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := c.orderService.CreateOrder(ctx.Request.Context(), userID.(string), &order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	orders, err := c.orderService.GetAllOrders(ctx.Request.Context(), userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
