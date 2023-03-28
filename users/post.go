package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"plusone/backend/database"
	"plusone/backend/utils"
)

func getLatestPost(c *gin.Context) {
	user, _ := utils.GetUser(c)
	post, found, err := database.GetLastestPost(user.ID)
	if !found && err == nil {
		c.JSON(404, gin.H{
			"status":  404,
			"message": fmt.Sprintf("User with ID %v not found!", user.ID),
		})
		return
	} else if !found && err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal server error!",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"post":   post,
	})
}
