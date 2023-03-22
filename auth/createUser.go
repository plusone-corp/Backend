package auth

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"plusone/backend/database"
	"time"
)

func createUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	displayName := c.PostForm("displayName")
	salt := GenerateRandomSalt(10)

	user, found, err := database.GetByUsername(username)
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

	user = &database.User{
		Username:    username,
		DisplayName: displayName,
		CreatedAt:   time.Now().Unix(),
		IconURL:     "https://plusone-corp.github.io/PlusOne/logo/adaptive-icon-1024.png",
		ID:          primitive.NewObjectID(),
		Email:       email,
		Credentials: database.Credentials{
			Password:      HashPassword(password, salt),
			Hash:          salt,
			RefreshToken:  "",
			LastRefreshed: time.Now().Unix(),
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
