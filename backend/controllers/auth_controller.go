package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peter6866/foodie/config"
	"github.com/peter6866/foodie/services"
)

type AuthController struct {
	userService *services.UserService
}

func NewAuthController(userService *services.UserService) *AuthController {
	return &AuthController{userService: userService}
}

func GoogleLogin(context *gin.Context) {
	var stateString string = config.AppConfig.GOOGLE_RANDOM_STATE
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL(stateString)

	// return the url
	context.JSON(http.StatusOK, gin.H{"url": url})
}

func (c *AuthController) AddToCart(context *gin.Context) {
	var requestBody struct {
		MenuItemID string `json:"menuItemID"`
	}

	if err := context.BindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctx := context.Request.Context()

	userID, _ := context.Get("userID")
	err := c.userService.AddToCart(ctx, userID.(string), requestBody.MenuItemID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Added to cart"})
}

func (c *AuthController) RemoveFromCart(context *gin.Context) {
	id := context.Param("id")
	ctx := context.Request.Context()

	userID, _ := context.Get("userID")
	err := c.userService.RemoveFromCart(ctx, userID.(string), id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove from cart"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Removed from cart"})

}

func (c *AuthController) GetCartItems(context *gin.Context) {
	ctx := context.Request.Context()

	userID, _ := context.Get("userID")
	cartItems, err := c.userService.GetCartItems(ctx, userID.(string))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart items"})
		return
	}

	context.JSON(http.StatusOK, cartItems)
}
