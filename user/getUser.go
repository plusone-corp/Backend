package user

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/database"
)

func getUserIdHandler(c *gin.Context) {
	id := c.Param("id")
	user, found, error := database.GetByID(id)
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
	user, found, error := database.GetByUsername(name)
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
	user, found, error := database.GetByEmail(email)
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
