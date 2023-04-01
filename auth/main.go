package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRouters(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.GET("/refresh", RefreshRoute)
	auth.POST("/login", LoginRoute)
	auth.POST("/register", createUser)
	auth.Use(JwtMiddleware())
	{
	}
}
