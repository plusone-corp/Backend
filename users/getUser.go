package users

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"plusone/backend/database"
	"plusone/backend/errorHandler"
	"plusone/backend/types"
	"plusone/backend/utils"
)

func getUserIdHandler(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid ID",
		})
		return
	}
	userData, found, error := database.GetUserByID(objId)

	user := types.UserSensored{Username: userData.Username, Avatar: userData.Avatar, DisplayName: userData.DisplayName, Description: userData.Description, Level: userData.Level}

	if error != nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	} else if !found {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidID)
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
	userData, found, error := database.GetUserByUsername(name)

	user := types.UserSensored{Username: userData.Username, Avatar: userData.Avatar, DisplayName: userData.DisplayName, Description: userData.Description, Level: userData.Level}

	if error != nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	} else if !found {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidID)
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
	userData, found, error := database.GetUserByEmail(email)

	user := types.UserSensored{Username: userData.Username, Avatar: userData.Avatar, DisplayName: userData.DisplayName, Description: userData.Description, Level: userData.Level}

	if error != nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	} else if !found {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidID)
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "User found!",
		"user":    user,
	})
}

func getMe(c *gin.Context) {
	user, claims := utils.GetUser(c)
	if user == nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"userID": claims.ID,
		"user":   user,
	})
}
