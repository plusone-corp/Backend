package auth

import (
	"github.com/gin-gonic/gin"
)

func invalidAuthorizationMethod(c *gin.Context) {
	c.JSON(400, gin.H{
		"status":  400,
		"message": "Invalid authorization method. Please use Bearer instead.",
	})
	return
}

func invalidToken(c *gin.Context) {
	c.JSON(401, gin.H{
		"status":  401,
		"message": "Invalid token, please check your credentials again!",
	})
}
