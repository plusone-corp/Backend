package utils

import (
	"github.com/gin-gonic/gin"
	"plusone/backend/database"
	"plusone/backend/types"
)

// GetUser Get userdata from JWT_PAYLOAD
func GetUser(c *gin.Context) (*types.UserResponse, *types.SignedDetails) {
	// get the data from key
	payload, exist := c.Get("JWT_PAYLOAD")
	if !exist {
		return nil, nil
	}

	// Parse the data to usable types
	claims := payload.(*types.SignedDetails)

	// Parse userid from string to primitive.ObjectID
	userId, err := StringToObjectId(claims.ID)

	// Get user data
	user, found, err := database.GetUserByID(*userId)
	if !found && err == nil {
		return nil, nil
	} else if !found && err != nil {
		return nil, nil
	}

	// Generate structs
	events := []types.Event{}
	friends := []types.UserFiltered{}

	// If there are events or friends data, get them
	if len(user.Events) > 0 {
		res, err := database.GetAllEvent(user.ID)
		if err != nil {
			return nil, nil
		}
		events = *res
	}
	if len(user.Friends) > 0 {
		res, err := database.GetAllFriends(user.ID)
		if !found && err != nil {
			return nil, nil
		}
		friends = *res
	}

	// Apply the data
	newUser := types.UserResponse{
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

	// Return the userdata and claim (this can be used in getMe() or something like that)
	return &newUser, claims
}
