package errorHandler

import (
	"github.com/gin-gonic/gin"
)

var (
	InvalidMethod            = "Invalid authorization method. Please use Bearer instead."
	AuthorizationKeyNotFound = "Authorization header not found, please contact support!"
	InvalidToken             = "Invalid token, please check your credentials again!"
	FailedTokenValidation    = "Failed to validate token, please login again!"
	InvalidFormBody          = "Invalid form body, please contact support!"
	InternalServerError      = "Internal server error, please try again later!"
	UsersNotFound            = "There are no users found."
	PostNotFound             = "No posts found for that user!"
	UsernameExisted          = "Username already existed"
	InvalidID                = "The document ID provided is invalid or the document with that id is not found!"
	RefreshToken             = "Access token expired, please refresh the token!"
)

func Unauthorized(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
	})
	return
}
