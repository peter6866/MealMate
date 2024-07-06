package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/peter6866/foodie/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	router.GET("/google_login", controllers.GoogleLogin)
	router.GET("/google_callback", controllers.GoogleCallBack)

	return router
}
