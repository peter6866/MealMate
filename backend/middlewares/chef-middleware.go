package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peter6866/foodie/config"
	"github.com/peter6866/foodie/repositories"
	"github.com/peter6866/foodie/services"
)

func ChefMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		client := config.MongoClient
		userRepo := repositories.NewUserRepository(client)
		userService := services.NewUserService(userRepo)
		userId, _ := ctx.Get("userId")

		user, err := userService.GetUser(ctx.Request.Context(), userId.(string))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if !user.IsChef {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		ctx.Next()
	}
}
