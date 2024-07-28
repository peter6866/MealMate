package controllers

import (
	"encoding/json"
	"net/http"
	"time"

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

func (c *MealController) UpdateMealFromOrder(ctx *gin.Context) {
	if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	mealID := ctx.Param("mealID")
	mealType := ctx.PostForm("mealType")
	withPartner := ctx.PostForm("withPartner")

	if mealID == "" || mealType == "" || withPartner == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing image file"})
		return
	}

	userID, _ := ctx.Get("userID")

	err = c.mealService.UpdateMealFromOrder(ctx.Request.Context(), userID.(string), mealID, *file, mealType, withPartner)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Meal updated successfully"})
}

func (c *MealController) GetAllMeals(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	meals, err := c.mealService.GetAllMeals(ctx.Request.Context(), userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"meals": meals})
}

// get meals by date range
func (c *MealController) GetMealsByDateRange(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	var input struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	meals, err := c.mealService.GetMealsByDateRange(ctx.Request.Context(), startDate, endDate, userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"meals": meals})
}

// delete a meal
func (c *MealController) DeleteMeal(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	mealID := ctx.Param("mealID")

	err := c.mealService.DeleteMeal(ctx.Request.Context(), userID.(string), mealID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Meal deleted successfully"})
}
