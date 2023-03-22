package user

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/auth"
)

func UserHandler(route *gin.RouterGroup) {
	group := route.Group("/user")
	group.POST("/create", createUser)
	group.Use(auth.AuthMiddleware.MiddlewareFunc())
	{
		group.GET("/getId/:id", getUserIdHandler)
		group.GET("/getName/:name", getUserNameHandler)
		group.GET("/getEmail/:email", getUserEmailHandler)
	}
}
