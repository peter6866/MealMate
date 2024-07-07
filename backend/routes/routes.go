package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/peter6866/foodie/controllers"
	"github.com/peter6866/foodie/middlewares"
	"github.com/peter6866/foodie/repositories"
	"github.com/peter6866/foodie/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	userRepo := repositories.NewUserRepository(client)
	userService := services.NewUserService(userRepo)
	authController := controllers.NewAuthController(userService)

	router.GET("/google_login", controllers.GoogleLogin)
	router.POST("/api/auth/loginOrRegister", authController.LoginOrRegister)

	menuItemRepo := repositories.NewMenuItemRepository(client)
	menuItemService := services.NewMenuItemService(menuItemRepo)
	menuItemController := controllers.NewMenuItemController(menuItemService)

	router.GET("/api/menuItems", menuItemController.GetAllMenuItems)

	// Authenticated routes
	authenticatedRoutes := router.Group("/api")
	authenticatedRoutes.Use(middlewares.AuthMiddleware)
	{
		menuItemRoutes := authenticatedRoutes.Group("/menuItems")
		{
			menuItemRoutes.GET("/:id", menuItemController.GetMenuItem)

			menuItemRoutes.POST("", middlewares.AdminMiddleware(), menuItemController.CreateMenuItem)
			// menuItemRoutes.PUT("/:id", middlewares.AdminMiddleware(), menuItemController.UpdateMenuItem)
			menuItemRoutes.DELETE("/:id", middlewares.AdminMiddleware(), menuItemController.DeleteMenuItem)
		}
	}

	return router
}
