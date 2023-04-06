package users

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"plusone/backend/database"
	"plusone/backend/errorHandler"
	"plusone/backend/types"
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
	userData, found, err := database.GetUserByID(objId)

	user := types.UserFiltered{Username: userData.Username, Avatar: userData.Avatar, DisplayName: userData.DisplayName, Description: userData.Description, Level: userData.Level}

	if err != nil {
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
	userData, found, err := database.GetUserByUsername(name)

	user := types.UserFiltered{Username: userData.Username, Avatar: userData.Avatar, DisplayName: userData.DisplayName, Description: userData.Description, Level: userData.Level}

	if err != nil {
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
	userData, found, err := database.GetUserByEmail(email)

	user := types.UserFiltered{Username: userData.Username, Avatar: userData.Avatar, DisplayName: userData.DisplayName, Description: userData.Description, Level: userData.Level}

	if err != nil {
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
