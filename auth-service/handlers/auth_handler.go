package handlers

import (
	"auth-service/config"
	"auth-service/services"
	"auth-service/utils"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func GoogleLogin(context *gin.Context) {
	var stateString string = config.AppConfig.GOOGLE_RANDOM_STATE
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL(stateString)

	// return the url
	context.JSON(http.StatusOK, gin.H{"url": url})
}

func (c *AuthHandler) LoginOrRegister(context *gin.Context) {
	var requestBody struct {
		Code  string `json:"code"`
		State string `json:"state"`
	}

	if err := context.BindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var stateString string = config.AppConfig.GOOGLE_RANDOM_STATE
	if requestBody.State != stateString {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter"})
		return
	}

	googlecon := config.AppConfig.GoogleLoginConfig
	token, err := googlecon.Exchange(context, requestBody.Code)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to exchange token"})
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info"})
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read response body"})
		return
	}

	var userInfo map[string]interface{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal response body"})
		return
	}

	ctx := context.Request.Context()

	// Find or create user
	user, err := c.userService.FindOrCreateUser(ctx, userInfo["name"].(string), userInfo["email"].(string), userInfo["id"].(string), "user", userInfo["picture"].(string))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find or create user"})
		return
	}

	jwtToken, err := utils.GenerateToken(user.ID, userInfo["email"].(string), user.Role)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	var isChef string
	if user.IsChef {
		isChef = "true"
	} else {
		isChef = "false"
	}

	context.JSON(http.StatusOK, gin.H{"token": jwtToken, "isChef": isChef})
}

func (c *AuthHandler) GetUser(context *gin.Context) {
	userID, _ := context.Get("userID")
	user, err := c.userService.GetUser(context.Request.Context(), userID.(string))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (c *AuthHandler) SetChefAndPartner(context *gin.Context) {
	var requestBody struct {
		IsChef       bool   `json:"isChef"`
		PartnerEmail string `json:"partnerEmail"`
	}

	if err := context.BindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctx := context.Request.Context()

	userID, _ := context.Get("userID")
	updatedUser, err := c.userService.SetChefAndPartner(ctx, userID.(string), requestBody.IsChef, requestBody.PartnerEmail)
	if err != nil {
		// if the err has a message, return the message
		if err.Error() != "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set chef and partner"})
	}

	context.JSON(http.StatusOK, updatedUser)
}

func (c *AuthHandler) AddToCart(context *gin.Context) {
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

func (c *AuthHandler) RemoveFromCart(context *gin.Context) {
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

// TODO: Implement GetCartItems
// func (c *AuthHandler) GetCartItems(context *gin.Context) {
// 	ctx := context.Request.Context()

// 	userID, _ := context.Get("userID")
// 	cartItems, err := c.userService.GetCartItems(ctx, userID.(string))
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart items"})
// 		return
// 	}

// 	context.JSON(http.StatusOK, cartItems)
// }
