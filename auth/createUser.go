package auth

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"plusone/backend/database"
	"plusone/backend/errorHandler"
	"plusone/backend/types"
	"time"
)

func createUser(c *gin.Context) {
	var userData types.UserCreate

	if c.ShouldBind(&userData) != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}
	salt := GenerateRandomSalt(10)

	user, found, err := database.GetUserByUsername(userData.Username)
	if err != nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	} else if found {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.UsernameExisted)
		return
	}

	user = &types.User{
		Username:    userData.Username,
		DisplayName: userData.DisplayName,
		CreatedAt:   time.Now(),
		Avatar:      "https://plusone-corp.github.io/PlusOne/logo/adaptive-icon-1024.png",
		ID:          primitive.NewObjectID(),
		Email:       userData.Email,
		Credentials: types.Credentials{
			Password:      HashPassword(userData.Password, salt),
			Hash:          salt,
			RefreshToken:  "",
			LastRefreshed: time.Now(),
		},
		Description: userData.Description,
		Age:         0,
		Friends:     []primitive.ObjectID{},
		Events:      []primitive.ObjectID{},
		Level: types.Level{
			Exp:    0,
			Level:  0,
			Badges: 0,
		},
	}

	ok, err := database.CreateUser(*user)
	if err != nil || !ok {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}

	c.JSON(202, gin.H{
		"status":  202,
		"message": "User created successfully",
	})
	return
}
