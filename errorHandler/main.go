package errorHandler

import (
	"github.com/gin-gonic/gin"
)

func Unauthorized(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
	})
	return
}
