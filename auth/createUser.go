package auth

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"plusone/backend/database"
	"plusone/backend/types"
	"time"
)

func createUser(c *gin.Context) {
	var userData types.UserCreate

	if c.ShouldBind(&userData) == nil {
		log.Println(userData.Username)
	}
	salt := GenerateRandomSalt(10)

	user, found, err := database.GetByUsername(userData.Username)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
		})
		return
	} else if found {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Username already existed",
		})
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
			LastRefreshed: time.Now(),
		},
	}

	ok, err := database.CreateUser(*user)
	if err != nil || !ok {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(202, gin.H{
		"status":  202,
		"message": "User created successfully",
	})
	return
}
