package auth

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/rand"
	"net/http"
	"plusone/backend/database"
	"plusone/backend/errorHandler"
	"plusone/backend/types"
	"plusone/backend/utils"
	"time"
)

// LoginRoute
/*
	 Input: types.Login
	Output: {
			"status": httpStatusCode,
			"token":  types.Token,
			"user":   types.User,
		}
*/
func LoginRoute(c *gin.Context) {
	var loginVals types.Login

	// Transforming the body request to json format with types.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, "Invalid login form fields")
		return
	}

	// Declaring username and password variable
	userID := loginVals.Username
	password := loginVals.Password

	// Validating username password, by calling the api and comparing the hash
	user, checkPassword, err := authorizeUser(userID, password)
	if err != nil {
		errorHandler.Unauthorized(c, http.StatusNotFound, fmt.Sprintf("User with username %v doesn't exist", userID))
		return
	}

	// If the credentials are true, then sign a token and return the tokens and userdata
	if checkPassword {
		tokens, err := Sign((*user).ID)
		if err != nil {
			errorHandler.Unauthorized(c, http.StatusInternalServerError, "Failed to sign an access token")
			return
		}

		c.JSON(200, gin.H{
			"status": 200,
			"token":  tokens,
			"user":   user,
		})
		return
	}

	// If it is false, return Unauthorized status code
	errorHandler.Unauthorized(c, http.StatusUnauthorized, "Username or password are invalid")
	return
}

// RefreshRoute
/*
	Input: Authorization header with Token "Bearer TOKEN"
	Output: {
		"status": httpStatusCode,
		"token": types.Token, (New access token and refresh token)
	}
*/
func RefreshRoute(c *gin.Context) {
	// Validating the headers and getting the refresh token from that
	// "Bearer TOKEN" -> "TOKEN"
	token, err := validateHeaders(c.GetHeader("Authorization"))
	if err != nil {
		// If token not found, return error
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.AuthorizationKeyNotFound)
		return
	}

	// Validating the token, checks if the token is valid
	claims, valid, err := ParseRefreshToken(*token)
	if err != nil || !valid {
		// If it is not valid, return error
		errorHandler.Unauthorized(c, http.StatusUnauthorized, *err)
		return
	}

	// Parsing the objectID (string) to be objectID (primitive.ObjectID)
	id, parseErr := utils.StringToObjectId(claims.ID)
	if parseErr != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, "Failed to validate user's id")
		return
	}

	// Getting user from the database, checks if the user exist
	user, found, userErr := database.GetUserByID(*id)
	if userErr != nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if !found {
		errorHandler.Unauthorized(c, http.StatusBadRequest, "The refresh token doesn't belong to any user!")
		return
	}

	// Calculate the next expiration date of the token
	calcTime := time.Now().Unix() - user.Credentials.LastRefreshed.Unix()
	// Refresh token, expiration date (7 days)
	week := int64(1000 * 60 * 60 * 24 * 7)
	// Access token, expiration date (1 hour)
	hour := int64(1000 * 60 * 60)

	// If the last refresh time is between 7 days and 1-hour period
	if calcTime < week && calcTime > hour {
		// If the refresh token that stored in database is the same as the request token, return error
		if user.Credentials.RefreshToken != *token {
			errorHandler.Unauthorized(c, http.StatusUnauthorized, "Invalid refresh token")
			return
		}
		// Else, sign a new token
		tokens, signErr := Sign(*id)
		if signErr != nil {
			errorHandler.Unauthorized(c, http.StatusInternalServerError, "Failed to sign an access token")
			return
		}

		// Return the token
		c.JSON(200, gin.H{
			"status": 200,
			"token":  tokens,
		})
		return
	} else if calcTime < hour {
		// If the last refresh time is less than 1-hour, return error, cause refresh too fast
		errorHandler.Unauthorized(c, http.StatusBadRequest, "You can't refresh the token, while the access token still valid!")
		return
	} else {
		// If the token is expired, return error
		errorHandler.Unauthorized(c, http.StatusBadRequest, "Refresh token already expired, please use login!")
		return
	}
}

// Logout
/*
	Input: Authorization header with Token "Bearer TOKEN"
	Output: {
		"status": httpStatusCode,
	}
*/
func Logout(c *gin.Context) {
	// Getting the token from header
	token, err := validateHeaders(c.GetHeader("Authorization"))
	if err != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.AuthorizationKeyNotFound)
		return
	}

	// Validating the access token
	claims, valid, err := ParseAccessToken(*token)
	if err != nil || !valid {
		errorHandler.Unauthorized(c, http.StatusForbidden, *err)
		return
	}

	// Validate the user id from string to primitive.ObjectID
	id, parseErr := utils.StringToObjectId(claims.ID)
	if parseErr != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, "Failed to validate user's id")
		return
	}

	// Remove the refresh token from database
	updateErr := database.RemoveRefreshToken(*id)
	if updateErr != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, "Failed to update user's id")
		return
	}

	// Blacklist the token
	LoggedOutToken.Set(claims.ID, *token, claims.ExpiresAt.Time)

	// Return code OK
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Logged out successfully",
	})
	return
}

// CreateUser
/*
	Input: types.UserForm
	Output: {
		"status": httpStatus,
		"message": string
	}
*/
func CreateUser(c *gin.Context) {
	var userData types.UserForm

	// Validating the request body to types.UserForm
	if c.ShouldBind(&userData) != nil {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}

	// If the request body doesn't match to the types, return error
	if !utils.UserFormValidation(userData) {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.InvalidFormBody)
		return
	}

	// Generate a hash, so we can hash the password
	salt := generateHashSalt(10)

	// Check if the username already exist
	_, found, err := database.GetUserByUsername(userData.Username)
	if err != nil {
		log.Println("Get user", err)
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	} else if found {
		errorHandler.Unauthorized(c, http.StatusBadRequest, errorHandler.UsernameExisted)
		return
	}

	// If display doesn't exist, then assign username to display
	displayName := userData.DisplayName
	if displayName == "" {
		displayName = userData.Username
	}

	// Create a full user types
	newUser := types.User{
		Username:    userData.Username,
		DisplayName: userData.DisplayName,
		CreatedAt:   time.Now(),
		// Default avatar
		Avatar:   "https://plusone-corp.github.io/PlusOne/logo/adaptive-icon-1024.png",
		ID:       primitive.NewObjectID(),
		Email:    userData.Email,
		Location: types.GeoJSON{},
		Credentials: types.Credentials{
			// Hash password here
			Password:      hashPassword(userData.Password, salt),
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

	// Create user
	ok, err := database.CreateUser(newUser)
	if err != nil || !ok {
		log.Println(err)
		errorHandler.Unauthorized(c, http.StatusInternalServerError, errorHandler.InternalServerError)
		return
	}

	tokens, err := Sign(newUser.ID)
	if err != nil {
		errorHandler.Unauthorized(c, http.StatusInternalServerError, "Failed to sign an access token")
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"token":  tokens,
		"user":   newUser,
	})
	return
}

// generateHashSalt
// Create a hash
func generateHashSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

// hashPassword
// hash the password with the generated hash
func hashPassword(password string, salt []byte) string {
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

// doPasswordsMatch
// match the password
func doPasswordsMatch(hashedPassword, currPassword string, salt []byte) bool {
	var currPasswordHash = hashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}

// authorizeUser
// check if the username and password are correct
func authorizeUser(username string, password string) (*types.User, bool, error) {
	user, found, err := database.GetUserByUsername(username)
	if err != nil {
		return nil, false, err
	} else if !found && err == nil {
		return nil, false, nil
	}
	isCorrect := doPasswordsMatch(user.Credentials.Password, password, []byte(user.Credentials.Hash))
	return user, isCorrect, nil
}
