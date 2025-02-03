package routes

import (
	"auth-service/config"
	"auth-service/handlers"
	"auth-service/middlewares"
	"auth-service/repositories"
	"auth-service/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	// config cors
	config := cors.Config{
		AllowOrigins:     []string{config.AppConfig.ALLOWED_ORIGIN},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))

	userRepo := repositories.NewUserRepository(client)
	userService := services.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(userService)

	router.GET("/google_login", handlers.GoogleLogin)
	router.POST("/api/auth/loginOrRegister", authHandler.LoginOrRegister)

	// Authenticated routes
	authenticatedRoutes := router.Group("/api")
	authenticatedRoutes.Use(middlewares.AuthMiddleware)
	{
		authenticatedRoutes.GET("/auth/getUser", authHandler.GetUser)
		authenticatedRoutes.POST("/auth/setChefAndPartner", authHandler.SetChefAndPartner)

		authenticatedRoutes.POST("/cart", authHandler.AddToCart)
		// authenticatedRoutes.GET("/cart", authHandler.GetCartItems)
		authenticatedRoutes.DELETE("/cart/:id", authHandler.RemoveFromCart)
	}

	return router
}
