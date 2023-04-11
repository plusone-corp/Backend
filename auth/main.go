package auth

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/utils"
)

var (
	LoggedOutToken = utils.NewBlockToken()
)

func AuthRouters(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.GET("/refresh", RefreshRoute)
	auth.POST("/login", LoginRoute)
	auth.POST("/register", CreateUser)
	auth.Use(JwtMiddleware())
	{
		auth.GET("/logout", Logout)
	}
}
