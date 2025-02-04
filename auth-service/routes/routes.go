package routes

import (
	"auth-service/config"
	"auth-service/handlers"
	"auth-service/middlewares"
	"auth-service/repositories"
	"auth-service/services"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var userService *services.UserService

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

	// Initialize services
	userRepo := repositories.NewUserRepository(client)
	userService = services.NewUserService(userRepo, config.RabbitMQChannel)
	authHandler := handlers.NewAuthHandler(userService)

	// Initialize and start the order events consumer
	consumer := services.NewOrderEventConsumer(userService, config.RabbitMQChannel)
	if err := consumer.Start(); err != nil {
		log.Printf("Warning: Failed to start order events consumer: %v", err)
	}

	// Routes setup
	router.GET("/api/auth/google_login", handlers.GoogleLogin)
	router.POST("/api/auth/loginOrRegister", authHandler.LoginOrRegister)

	// Authenticated routes
	authenticatedRoutes := router.Group("/api")
	authenticatedRoutes.Use(middlewares.AuthMiddleware)
	{
		authenticatedRoutes.GET("/auth/getUser", authHandler.GetUser)
		authenticatedRoutes.POST("/auth/setChefAndPartner", authHandler.SetChefAndPartner)

		authenticatedRoutes.POST("/auth/cart", authHandler.AddToCart)
		// authenticatedRoutes.GET("/cart", authHandler.GetCartItems)
		authenticatedRoutes.DELETE("/auth/cart/:id", authHandler.RemoveFromCart)
	}

	return router
}
