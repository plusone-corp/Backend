package utils

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"plusone/backend/config"
	"plusone/backend/database"
	"plusone/backend/types"
)

func GetUser(c *gin.Context) (*types.ResUser, jwt.MapClaims) {
	claims := jwt.ExtractClaims(c)
	username, _ := c.Get(config.IDENTIFY_KEY)
	user, found, err := database.GetByUsername(username.(*types.User).Username)
	if !found && err == nil {
		return nil, nil
	} else if !found && err != nil {
		return nil, nil
	}

	events := []types.Event{}
	friends := []types.UserSensored{}

	if len(user.Events) > 0 {
		res, found, err := database.GetManyEventsID(user.Events)
		if !found && err != nil {
			log.Println("Events ", err)
			return nil, nil
		}
		events = *res
	}
	if len(user.Friends) > 0 {
		res, found, err := database.GetManyUserID(user.Friends)
		if !found && err != nil {
			log.Println("Friends ", err)
			return nil, nil
		}
		friends = *res
	}

	newUser := types.ResUser{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Avatar:      user.Avatar,
		DisplayName: user.DisplayName,
		Description: user.Description,
		Age:         user.Age,
		CreatedAt:   user.CreatedAt,
		Events:      events,
		Friends:     friends,
		Location:    user.Location,
		Level:       user.Level,
	}

	return &newUser, claims
}
