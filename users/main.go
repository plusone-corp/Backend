package users

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/auth"
)

func UserHandler(route *gin.Engine) {
	group := route.Group("/users")
	getUser(group)
	group.Use(auth.AuthMiddleware.MiddlewareFunc())
	{
		group.GET("/getId/:id", getUserIdHandler)
		group.GET("/getName/:name", getUserNameHandler)
		group.GET("/getEmail/:email", getUserEmailHandler)
	}
}
