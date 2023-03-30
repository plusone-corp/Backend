package main

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/auth"
	"plusone/backend/events"
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

	auth.AuthRouters(Router)
	users.UserHandler(Router)
	posts.PostHandlers(Router)
	events.EventsHandlers(Router)

	err := Router.Run(":80")
	if err != nil {
		panic(err)
	}
}

func hello(c *gin.Context) {
	c.Status(200)
}

func handleNotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"status":  404,
		"message": "Page not found!",
	})
}
