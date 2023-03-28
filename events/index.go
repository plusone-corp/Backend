package events

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/auth"
)

func EventsHandlers(c *gin.Engine) {
	group := c.Group("/events")
	group.Use(auth.AuthMiddleware.MiddlewareFunc())
	{
		group.GET("/get/:id", getEventID)
		group.POST("/create", createEvent)
	}
}
