package routes

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/peter6866/foodie/config"
	"github.com/peter6866/foodie/controllers"
	"github.com/peter6866/foodie/middlewares"
	"github.com/peter6866/foodie/repositories"
	"github.com/peter6866/foodie/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	// config cors
	corsConfig := cors.Config{
		AllowOrigins:     []string{config.AppConfig.ALLOWED_ORIGIN},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))
	// Register custom error handler middleware
	router.Use(middlewares.ErrorHandler())

	// All Repositories
	userRepo := repositories.NewUserRepository(client)
	simpUserRepo := repositories.NewSimpUserRepository(client)
	categoryRepo := repositories.NewCategoryRepository(client)
	menuItemRepo := repositories.NewMenuItemRepository(client)
	orderRepo := repositories.NewOrderRepository(client)
	mealRepo := repositories.NewMealRepository(client)

	consumer := services.NewUserEventConsumer(simpUserRepo, config.RabbitMQChannel)
	if err := consumer.Start(); err != nil {
		log.Printf("Warning: Failed to start consumer: %v", err)
	}

	// All Services
	userService := services.NewUserService(userRepo, menuItemRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	menuItemService := services.NewMenuItemService(menuItemRepo, categoryRepo, simpUserRepo)
	orderService := services.NewOrderService(userRepo, orderRepo, mealRepo)
	mealService := services.NewMealService(mealRepo)

	authController := controllers.NewAuthController(userService)

	categoryController := controllers.NewCategoryController(categoryService)
	menuItemController := controllers.NewMenuItemController(menuItemService)
	orderController := controllers.NewOrderController(orderService)
	mealController := controllers.NewMealController(mealService)

	router.GET("/api/categories", categoryController.GetAllCategories)

	// Authenticated routes
	authenticatedRoutes := router.Group("/api")
	authenticatedRoutes.Use(middlewares.AuthMiddleware)
	{
		authenticatedRoutes.GET("/cart", authController.GetCartItems)
		authenticatedRoutes.DELETE("/cart/:id", authController.RemoveFromCart)

		menuItemRoutes := authenticatedRoutes.Group("/menuItems")
		{
			menuItemRoutes.GET("/:id", menuItemController.GetMenuItem)

			menuItemRoutes.GET("", menuItemController.GetAllMenuItems)
			menuItemRoutes.POST("", middlewares.ChefMiddleware(), menuItemController.CreateMenuItem)
			// menuItemRoutes.PUT("/:id", middlewares.ChefMiddleware(), menuItemController.UpdateMenuItem)
			menuItemRoutes.DELETE("/:id", middlewares.ChefMiddleware(), menuItemController.DeleteMenuItem)
		}

		categoryRoutes := authenticatedRoutes.Group("/categories")
		{
			categoryRoutes.POST("", middlewares.ChefMiddleware(), categoryController.CreateCategory)
		}

		orderRoutes := authenticatedRoutes.Group("/orders")
		{
			orderRoutes.POST("", orderController.CreateOrder)
			orderRoutes.GET("", orderController.GetAllOrders)
			orderRoutes.PUT("/:id", orderController.CompleteOrder)
		}

		mealRoutes := authenticatedRoutes.Group("/meals")
		{
			mealRoutes.POST("", mealController.CreateMeal)
			mealRoutes.PUT("/:mealID", mealController.UpdateMealFromOrder)
			mealRoutes.GET("", mealController.GetAllMeals)
			mealRoutes.POST("/date", mealController.GetMealsByDateRange)
			mealRoutes.DELETE("/:mealID", mealController.DeleteMeal)
		}
	}

	return router
}
