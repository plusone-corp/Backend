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
		postGroup := group.Group("/post")
		postGroup.GET("/latest", getLatestPost)
		postGroup.GET("/all", getAllPost)
	}
}
