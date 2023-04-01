package utils

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/database"
	"plusone/backend/types"
)

func GetUser(c *gin.Context) (*types.ResUser, *types.SignedDetails) {
	payload, exist := c.Get("JWT_PAYLOAD")
	if !exist {
		return nil, nil
	}
	claims := payload.(*types.SignedDetails)
	userId, err := StringToObjectId(claims.ID)
	user, found, err := database.GetUserByID(*userId)
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
			return nil, nil
		}
		events = *res
	}
	if len(user.Friends) > 0 {
		res, found, err := database.GetManyUserID(user.Friends)
		if !found && err != nil {
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
