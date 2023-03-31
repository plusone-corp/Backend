package auth

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"net/http"
	"plusone/backend/database"
	"plusone/backend/errorHandler"
	"plusone/backend/types"
	"plusone/backend/utils"
	"time"
)

func AuthUser(username string, password string) (*primitive.ObjectID, bool, error) {
	user, found, err := database.GetUserByUsername(username)
	if err != nil {
		return nil, false, err
	} else if !found && err == nil {
		return nil, false, nil
	}
	isCorrect := doPasswordsMatch(user.Credentials.Password, password, []byte(user.Credentials.Hash))
	id := user.ID
	return &id, isCorrect, nil
}

func LoginRoute(c *gin.Context) {
	var loginVals types.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, "Invalid login form fields")
		return
	}
	userID := loginVals.Username
	password := loginVals.Password

	id, checkPassword, err := AuthUser(userID, password)
	if err != nil {
		errorHandler.Unauthorized(c, http.StatusNotFound, fmt.Sprintf("User with username %v doesn't exist", userID))
		return
	}

	if checkPassword {
		tokens, err := Sign(*id)
		if err != nil {
			errorHandler.Unauthorized(c, http.StatusInternalServerError, "Failed to sign an access token")
			return
		}

		c.JSON(200, gin.H{
			"status": 200,
			"token":  tokens,
		})
		return
	}

	errorHandler.Unauthorized(c, http.StatusUnauthorized, "Username or password are invalid")
	return
}

func RefreshRoute(c *gin.Context) {
	token := c.GetHeader("X-Token")

	claims, valid, err := ParseRefreshToken(token)
	if err != nil || !valid {
		errorHandler.Unauthorized(c, http.StatusUnauthorized, *err)
		return
	}

	id, parseErr := utils.StringToObjectId(claims.ID)
	if parseErr != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, "Failed to validate user's id")
		return
	}

	user, found, userErr := database.GetUserByID(*id)
	if userErr != nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if !found {
		errorHandler.Unauthorized(c, http.StatusBadRequest, "The refresh token doesn't belong to any user!")
		return
	}

	calcTime := time.Now().Unix() - user.Credentials.LastRefreshed.Unix()
	week := int64(1000 * 60 * 60 * 24 * 7)
	hour := int64(1000 * 60 * 60)

	if calcTime < week && calcTime > hour {
		if user.Credentials.RefreshToken != token {
			errorHandler.Unauthorized(c, http.StatusUnauthorized, "Invalid refresh token")
			return
		}
		tokens, signErr := Sign(*id)
		if signErr != nil {
			errorHandler.Unauthorized(c, http.StatusInternalServerError, "Failed to sign an access token")
			return
		}

		c.JSON(200, gin.H{
			"status": 200,
			"token":  tokens,
		})
		return
	} else if calcTime < hour {
		errorHandler.Unauthorized(c, http.StatusBadRequest, "You can't refresh the token, while the access token still valid!")
		return
	} else {
		errorHandler.Unauthorized(c, http.StatusBadRequest, "Refresh token already expired, please use login!")
		return
	}

}

func GenerateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func HashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

func doPasswordsMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = HashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}
