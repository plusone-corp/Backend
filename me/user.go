package me

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"plusone/backend/errorHandler"
	"plusone/backend/utils"
)

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
