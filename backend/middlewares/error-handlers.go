package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	custom_errors "github.com/peter6866/foodie/custom-errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			var status int
			var message string

			switch err {
			case custom_errors.ErrUnauthorized:
				status = http.StatusUnauthorized
				message = "You are not authorized"
			case custom_errors.ErrOrderNotFound:
				status = http.StatusNotFound
				message = "Order not found"
			case custom_errors.ErrOrderCompleted:
				status = http.StatusBadRequest
				message = "Order already completed"
			case custom_errors.ErrInvalidObjectID:
				status = http.StatusBadRequest
				message = "Invalid Object ID"
			default:
				status = http.StatusInternalServerError
				message = "Internal server error"
			}

			c.JSON(status, gin.H{"error": message})
		}
	}
}
