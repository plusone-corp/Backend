package auth

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

func createUser(c *gin.Context) {
	var userData types.UserForm

	if c.ShouldBind(&userData) != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}

	if !utils.UserFormValidation(userData) {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}

	salt := GenerateRandomSalt(10)

	_, found, err := database.GetUserByUsername(userData.Username)
	if err != nil {
		log.Println("Get user", err)
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	} else if found {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.UsernameExisted)
		return
	}

	displayName := userData.DisplayName
	if displayName == "" {
		displayName = userData.Username
	}

	newUser := types.User{
		Username:    userData.Username,
		DisplayName: userData.DisplayName,
		CreatedAt:   time.Now(),
		Avatar:      "https://plusone-corp.github.io/PlusOne/logo/adaptive-icon-1024.png",
		ID:          primitive.NewObjectID(),
		Email:       userData.Email,
		Location:    types.GeoJSON{},
		Credentials: types.Credentials{
			Password:      HashPassword(userData.Password, salt),
			Hash:          salt,
			RefreshToken:  "",
			LastRefreshed: time.Now(),
		},
		Description: "",
		Age:         userData.Age,
		Friends:     []primitive.ObjectID{},
		Events:      []primitive.ObjectID{},
		Level: types.Level{
			Exp:    0,
			Level:  0,
			Badges: 0,
		},
		Setting: types.Setting{
			Privacy: types.Privacy{
				ShareLocation:      true,
				AllowFriendRequest: true,
				AllowInvite:        true,
			},
		},
	}

	ok, err := database.CreateUser(newUser)
	if err != nil || !ok {
		log.Println(err)
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}

	c.JSON(202, gin.H{
		"status":  202,
		"message": "User created successfully",
	})
	return
}
