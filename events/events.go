package events

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

func getEventID(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}

	event, found, err := database.GetEventID(objId)
	if !found {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidID)
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
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
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
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}

	if err != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidID)
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
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Successfully created new event",
		"event":   newPost,
	})
}
