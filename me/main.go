package me

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/auth"
)

func MeHandlers(c *gin.Engine) {
	group := c.Group("/@me")
	group.Use(auth.JwtMiddleware())
	{
		group.GET("/", getMe)
	}
}
