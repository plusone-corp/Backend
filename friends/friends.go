package friends

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"plusone/backend/database"
	"plusone/backend/errorHandler"
	"plusone/backend/types"
	"plusone/backend/utils"
)

func getFriendList(c *gin.Context) {
	user, _ := utils.GetUser(c)
	friends, err := database.GetFriends(user.ID)
	if err != nil {
		log.Println(err)
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}

	log.Println(friends)

	if friends == nil || len(*friends) == 0 {
		errorHandler.Unauthorized(c, http.StatusNotFound, errorHandler.UsersNotFound)
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"users":  friends,
	})
	return
}

func addFriend(c *gin.Context) {
	user, _ := utils.GetUser(c)
	id := c.Param("id")
	err := database.AddFriend(user.ID, id)
	if err != nil {
		log.Println(err)
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Successfully updated user",
	})
	return
}

func getAllFriends(c *gin.Context) {
	user, _ := utils.GetUser(c)

	friends, err := database.GetAllFriends(user.ID)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  200,
			"friends": []*types.UserFiltered{},
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"friends": friends,
	})
}
