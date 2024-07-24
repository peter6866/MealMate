package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/services"
)

type MealController struct {
	mealService *services.MealService
}

func NewMealController(mealService *services.MealService) *MealController {
	return &MealController{mealService: mealService}
}

func (c *MealController) CreateMeal(ctx *gin.Context) {
	if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	jsonData := ctx.PostForm("json")
	if jsonData == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Empty JSON data"})
		return
	}

	var meal models.Meal
	if err := json.Unmarshal([]byte(jsonData), &meal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing image file"})
		return
	}

	userID, _ := ctx.Get("userID")

	err = c.mealService.CreateMeal(ctx.Request.Context(), userID.(string), &meal, *file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Meal created successfully"})
}
