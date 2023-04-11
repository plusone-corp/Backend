package main

import (
	"context"
	"plusone/backend/auth"
	"plusone/backend/config"
	"plusone/backend/events"
	"plusone/backend/posts"
	ratelimiter "plusone/backend/rateLimiter"
	"plusone/backend/types"
	"plusone/backend/users"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var (
	Router *gin.Engine
)

func main() {
	// Init
	Router = gin.Default()

	// Initialize context and Redis client	
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:	  config.REDIS_URL,
		Password: config.REDIS_SECRET,
	})


	// Attach ratelimiter middleware to main router
	Router.Use(ratelimiter.LimitRequest(ctx, rdb, []types.RateLimit{
		// ! Please note that include sub routes option may rate limit the non existant routes
		{ Route: "/auth", IncludeSubRoutes: true, RequestPerHour: 15 },
		{ Route: "/users", IncludeSubRoutes: true, RequestPerHour: 10 },
		{ Route: "/events", IncludeSubRoutes: true, RequestPerHour: 150 },
	}))

	Router.GET("/", hello)
	Router.NoRoute(handleNotFound)

	Router.Use(CorsMiddleware())

	auth.AuthRouters(Router)
	users.UserHandler(Router)
	posts.PostHandlers(Router)
	events.EventsHandlers(Router)

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
