package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peter6866/foodie/config"
	"github.com/peter6866/foodie/services"
)

type AuthController struct {
	userService *services.UserService
}

func GoogleLogin(context *gin.Context) {
	var stateString string = config.AppConfig.GOOGLE_RANDOM_STATE
	// Get the URL to redirect the user to
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL(stateString)
	// Redirect the user to the Google login page
	context.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallBack(context *gin.Context) {
	var stateString string = config.AppConfig.GOOGLE_RANDOM_STATE
	state := context.Query("state")
	if state != stateString {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter"})
		return
	}

	code := context.Query("code")

	googlecon := config.AppConfig.GoogleLoginConfig

	token, err := googlecon.Exchange(context, code)
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

	var user map[string]interface{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal response body"})
		return
	}

	context.JSON(http.StatusOK, user)
}
