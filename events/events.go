package events

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

func getEventID(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid ID",
		})
		return
	}

	event, found, err := database.GetEventID(objId)
	if !found {
		c.JSON(404, gin.H{
			"status":  404,
			"message": fmt.Sprintf("Event with ID %v not found!", objId),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"event":  event,
	})
}

func createEvent(c *gin.Context) {
	var form *types.EventCreate

	err := c.ShouldBind(&form)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid form body!",
		})
		return
	}

	user, _ := utils.GetUser(c)

	var ageLimit int
	if form.AgeLimit == nil {
		ageLimit = user.Age/2 + 7
	} else {
		ageLimit = *form.AgeLimit
	}

	invites, err := utils.StringToObjectIDs(form.Invites)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid form body!",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid ID",
		})
		return
	}

	post := types.Event{
		Id:          primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		Title:       form.Title,
		Description: form.Description,
		AgeLimit:    ageLimit,
		Author:      user.ID,
		Invites:     invites,
		Comments:    []types.Comment{},
		Reactions:   []types.Reaction{},
		Posts:       []primitive.ObjectID{},
	}

	newPost, err := database.CreateEvent(post)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal server error!",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Successfully created new event",
		"event":   newPost,
	})
}
