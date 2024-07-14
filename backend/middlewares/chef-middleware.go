package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChefMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, _ := ctx.Get("role")
		if role != "chef" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		ctx.Next()
	}
}
