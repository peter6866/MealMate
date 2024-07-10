package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peter6866/foodie/config"
	"github.com/peter6866/foodie/services"
	"github.com/peter6866/foodie/utils"
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

func (c *AuthController) LoginOrRegister(context *gin.Context) {
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

	context.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

func (c *AuthController) GetUser(context *gin.Context) {
	userID, _ := context.Get("userID")
	user, err := c.userService.GetUser(userID.(string))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	context.JSON(http.StatusOK, user)
}
