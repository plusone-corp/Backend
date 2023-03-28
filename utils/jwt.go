package utils

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"plusone/backend/config"
	"plusone/backend/database"
	"plusone/backend/types"
)

func GetUser(c *gin.Context) (*types.User, jwt.MapClaims) {
	claims := jwt.ExtractClaims(c)
	username, _ := c.Get(config.IDENTIFY_KEY)
	user, found, error := database.GetByUsername(username.(*types.User).Username)
	if error != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
		})
		return nil, nil
	} else if !found {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
		})
		return nil, nil
	}
	return user, claims
}
