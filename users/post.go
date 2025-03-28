package users

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"plusone/backend/database"
	"plusone/backend/errorHandler"
	"plusone/backend/utils"
)

func getLatestPost(c *gin.Context) {
	user, _ := utils.GetUser(c)
	post, found, err := database.GetLatestPost(user.ID)
	if !found && err == nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidID)
		return
	} else if !found && err != nil {
		log.Println(err)
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}

	if post == nil {
		errorHandler.Unauthorized(c, http.StatusNotFound, errorHandler.PostNotFound)
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"post":   post,
	})
}

func getAllPost(c *gin.Context) {
	user, _ := utils.GetUser(c)

	posts, found, err := database.GetAllPost(user.ID)
	if !found && err == nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidID)
		return
	} else if !found && err != nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"posts":  posts,
	})
}
