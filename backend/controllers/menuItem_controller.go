package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuItemController struct {
	menuItemService *services.MenuItemService
}

func NewMenuItemController(menuItemService *services.MenuItemService) *MenuItemController {
	return &MenuItemController{menuItemService: menuItemService}
}

func (c *MenuItemController) CreateMenuItem(ctx *gin.Context) {
	var menuItem models.MenuItem

	if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	categoryIdStr := ctx.Request.FormValue("categoryId")
	if categoryIdStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Category ID is required"})
		return
	}

	categoryId, err := primitive.ObjectIDFromHex(categoryIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	jsonData := ctx.Request.FormValue("json")

	if err := json.Unmarshal([]byte(jsonData), &menuItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing image file"})
		return
	}

	userID, _ := ctx.Get("userID")
	menuItem.CategoryID = categoryId

	err = c.menuItemService.CreateMenuItem(ctx.Request.Context(), userID.(string), &menuItem, *file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, menuItem)
}

func (c *MenuItemController) GetMenuItem(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, _ := ctx.Get("userID")
	menuItem, err := c.menuItemService.GetMenuItem(ctx.Request.Context(), id, userID.(string))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if menuItem == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "menu item not found"})
		return
	}

	ctx.JSON(http.StatusOK, menuItem)
}

func (c *MenuItemController) DeleteMenuItem(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, _ := ctx.Get("userID")

	err := c.menuItemService.DeleteMenuItem(ctx.Request.Context(), id, userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Menu item deleted successfully"})
}

func (c *MenuItemController) GetAllMenuItems(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	menuItems, err := c.menuItemService.GetAllMenuItems(ctx.Request.Context(), userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menuItems)
}

func (c *MenuItemController) GetMenuItemsByCategory(ctx *gin.Context) {
	categoryID := ctx.Query("category")

	menuItems, err := c.menuItemService.GetMenuItemsByCategory(ctx.Request.Context(), categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menuItems)
}

func (c *MenuItemController) GetMenuItemsBySpiceLevel(ctx *gin.Context) {
	spiceLevel := ctx.Query("spiceLevel")

	menuItems, err := c.menuItemService.GetMenuItemsBySpiceLevel(ctx.Request.Context(), models.SpiceLevel(spiceLevel))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menuItems)
}

func (c *MenuItemController) GetMenuItemsByAlcoholContent(ctx *gin.Context) {
	alcoholContent := ctx.Query("alcoholContent")

	menuItems, err := c.menuItemService.GetMenuItemsByAlcoholContent(ctx.Request.Context(), models.AlcoholContent(alcoholContent))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menuItems)
}
