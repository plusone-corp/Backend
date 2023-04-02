package main

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/auth"
	"plusone/backend/events"
	"plusone/backend/friends"
	"plusone/backend/posts"
	"plusone/backend/users"
)

var (
	Router *gin.Engine
)

func main() {
	// Init
	Router = gin.Default()
	Router.GET("/", hello)
	Router.NoRoute(handleNotFound)

	Router.Use(CorsMiddleware())

	auth.AuthRouters(Router)
	users.UserHandler(Router)
	posts.PostHandlers(Router)
	events.EventsHandlers(Router)
	friends.FriendsHandlers(Router)

	err := Router.Run(":80")
	if err != nil {
		panic(err)
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func hello(c *gin.Context) {
	c.Status(204)
}

func handleNotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"status":  404,
		"message": "Page not found!",
	})
}
