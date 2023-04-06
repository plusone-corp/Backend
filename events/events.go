package events

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
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
	var form types.EventForm

	err := c.ShouldBind(&form)
	if err != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}

	if !utils.EventFormValidation(form) {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}

	user, _ := utils.GetUser(c)

	log.Println("Age", user)

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

	geoloc, err := utils.StringArrToGeoJSON(form.Location)
	if err != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}

	if geoloc == nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}

	post := types.Event{
		Id:          primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		Title:       form.Title,
		Description: form.Description,
		AgeLimit:    ageLimit,
		Location:    *geoloc,
		Image:       form.Image,
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

func getLatestEvent(c *gin.Context) {
	user, _ := utils.GetUser(c)
	event, found, err := database.GetLatestEvent(user.ID)
	if !found && err == nil {
		errorHandler.Unauthorized(c, http.StatusNotFound, errorHandler.InvalidID)
		return
	} else if !found && err != nil {
		log.Println(err)
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}

	if event == nil {
		errorHandler.Unauthorized(c, http.StatusNotFound, errorHandler.PostNotFound)
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"event":  event,
	})
}

func getALlEvent(c *gin.Context) {
	user, _ := utils.GetUser(c)

	events, err := database.GetAllEvent(user.ID)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 200,
			"events": []*types.Event{},
		})
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"events": events,
	})
}
