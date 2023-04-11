package rateLimiter

import (
	"context"
	"plusone/backend/types"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func LimitRequest(ctx context.Context, rdb *redis.Client, routes []types.RateLimit) gin.HandlerFunc  {
	return func(c *gin.Context) {
		var targetRoute types.RateLimit
		for _, route := range routes {
			if route.IncludeSubRoutes && strings.HasPrefix(c.Request.URL.Path, route.Route) {
				targetRoute = route
			} else if !route.IncludeSubRoutes && c.Request.URL.Path == route.Route {
				targetRoute = route
			}
		}		

		// Route is not entered so it means that route is not rate limited
		if targetRoute.Route == "" {
			c.Next()

			return
		}

		rateLimitKey := c.Request.URL.Path + ":" + c.ClientIP()

		// Fetch client's reqeust count from Redis
		val, err := rdb.Get(ctx, rateLimitKey).Result()

		// If the request count registry does not exist then set the value to 0
		if err == redis.Nil {
			val = "0"
		}

		// Fetch the remaining time needed to expire the key
		ttl, _ := rdb.TTL(ctx, rateLimitKey).Result()

		// If expiration query not exists create a new one
		if ttl <= -1 * time.Nanosecond {
			rdb.Expire(ctx, rateLimitKey, time.Hour)
			ttl = time.Hour
		}

		// Convert request count value to number because Redis returns string 
		requestAmount, _ := strconv.ParseInt(val, 10, 64)

		// Set the remaining request amount to show it on response header
		remainingRequest := strconv.Itoa(int(targetRoute.RequestPerHour) - int(requestAmount + 1))
		if targetRoute.RequestPerHour <= (requestAmount + 1) { remainingRequest = "0" }

		// Write required headers to response
		c.Header("X-RateLimit-Limit", strconv.Itoa(int(targetRoute.RequestPerHour)))
		c.Header("X-RateLimit-Remaining", remainingRequest)
		c.Header("X-RateLimit-Reset", strconv.Itoa(int(ttl.Seconds())))
		c.Header("Retry-After", strconv.Itoa(int(ttl.Seconds())))

		// If the request amount exceeded/hit the max request amount then apply the procecure
		if requestAmount >= targetRoute.RequestPerHour {
			// Abort the request with a JSON object
			c.AbortWithStatusJSON(429, gin.H{
				"status": 429,
				"message": "Too many requests",
			})

		// If the user did not exceed the request amount per minute then increase user request amount and let the request complete 
		} else {
			// Increase user request amount and set thte value to redis
			rdb.Do(ctx, "SET", rateLimitKey, requestAmount + 1, "KEEPTTL")

			// Let the request continue
			c.Next()
		}

	}
}