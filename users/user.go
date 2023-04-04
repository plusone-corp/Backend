package users

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/auth"
)

func getUser(c *gin.RouterGroup) {
	group := c.Group("/@me")
	group.Use(auth.JwtMiddleware())
	{
		group.GET("/", getMe)
	}
}
