package friends

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/auth"
)

func FriendsHandlers(c *gin.Engine) {
	group := c.Group("/friends")
	group.Use(auth.JwtMiddleware())
	{
		group.GET("/list", getFriendList)
		group.GET("/add/:id", addFriend)
	}
}
