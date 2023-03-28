package posts

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"plusone/backend/database"
	"plusone/backend/types"
	"plusone/backend/utils"
	"time"
)

func getPostID(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid ID",
		})
		return
	}

	post, found, err := database.GetPostID(objId)
	if !found && err == nil {
		c.JSON(404, gin.H{
			"status":  404,
			"message": fmt.Sprintf("Post with ID %v not found!", id),
		})
		return
	} else if !found && err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal server error",
		})
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
		log.Println(err)
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid form body!",
		})
		return
	}

	user, _ := utils.GetUser(c)

	eventObjId, err := primitive.ObjectIDFromHex(post.Event)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid ID",
		})
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

	res, found, error := database.CreatePost(newPost)
	if !found && error == nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid form body! User not found!",
		})
		return
	} else if !found && error != nil {
		log.Println(error)
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal server error!",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Created post successfully",
		"post":    res,
	})
}
