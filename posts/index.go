package posts

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/auth"
)

func PostHandlers(c *gin.Engine) {
	group := c.Group("/posts")
	group.Use(auth.AuthMiddleware.MiddlewareFunc())
	{
		group.GET("/get/:id", getPostID)
		group.POST("/create", createPost)
	}
}
