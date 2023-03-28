package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	AuthMiddleware *jwt.GinJWTMiddleware
)

func AuthRouters(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.GET("/refresh", AuthMiddleware.RefreshHandler)
	auth.POST("/login", AuthMiddleware.LoginHandler)
	auth.POST("/register", createUser)
	auth.Use(AuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/logout", AuthMiddleware.LogoutHandler)
	}
}
