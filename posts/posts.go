package posts

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"plusone/backend/database"
	"plusone/backend/errorHandler"
	"plusone/backend/types"
	"plusone/backend/utils"
	"time"
)

func getPostID(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}

	post, found, err := database.GetPostID(objId)
	if !found && err == nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidID)
		return
	} else if !found && err != nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Successfully fetched a post with ID " + id,
		"data":    post,
	})
}

func createPost(c *gin.Context) {
	var post *types.PostCreate

	err := c.ShouldBind(&post)
	if err != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}

	user, _ := utils.GetUser(c)

	eventObjId, err := primitive.ObjectIDFromHex(post.Event)
	if err != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidID)
		return
	}

	newPost := types.Post{
		Id:          primitive.NewObjectID(),
		Description: post.Description,
		Title:       post.Title,
		Author:      user.ID,
		Event:       eventObjId,
		Image:       post.Image,
		Reactions:   []types.Reaction{},
		Comments:    []types.Comment{},
		CreatedAt:   time.Now(),
	}

	res, found, err := database.CreatePost(newPost)
	if !found && err == nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidID)
		return
	} else if !found && err != nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Created post successfully",
		"post":    res,
	})
}
