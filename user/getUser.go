package user

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/database"
	"plusone/backend/types"
)

func getUserIdHandler(c *gin.Context) {
	id := c.Param("id")
	userData, found, error := database.GetByID(id)

	user := types.UserSensored{Username: userData.Username, Avatar: userData.Avatar, DisplayName: userData.DisplayName, Description: userData.Description, Events: userData.Events, Posts: userData.Posts, Level: userData.Level}

	if error != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error!",
		})
		return
	} else if !found {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "User ID not found!",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "User found!",
		"user":    user,
	})
}

func getUserNameHandler(c *gin.Context) {
	name := c.Param("name")
	userData, found, error := database.GetByUsername(name)

	user := types.UserSensored{Username: userData.Username, Avatar: userData.Avatar, DisplayName: userData.DisplayName, Description: userData.Description, Events: userData.Events, Posts: userData.Posts, Level: userData.Level}

	if error != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error!",
		})
		return
	} else if !found {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "User ID not found!",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "User found!",
		"user":    user,
	})
}

func getUserEmailHandler(c *gin.Context) {
	email := c.Param("email")
	userData, found, error := database.GetByEmail(email)

	user := types.UserSensored{Username: userData.Username, Avatar: userData.Avatar, DisplayName: userData.DisplayName, Description: userData.Description, Events: userData.Events, Posts: userData.Posts, Level: userData.Level}

	if error != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error!",
		})
		return
	} else if !found {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "User ID not found!",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "User found!",
		"user":    user,
	})
}
