package auth

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"plusone/backend/config"
	"plusone/backend/database"
)

var (
	AuthMiddleware *jwt.GinJWTMiddleware
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	username, _ := c.Get(config.IDENTIFY_KEY)
	user, found, error := database.GetByUsername(username.(*database.User).Username)
	if error != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
		})
		return
	} else if !found {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"userID": claims[config.IDENTIFY_KEY],
		"user": database.SensoredUser{
			Username:    user.Username,
			DisplayName: user.DisplayName,
			CreatedAt:   user.CreatedAt,
			IconURL:     user.IconURL,
			ID:          user.ID,
			Email:       user.Email,
			Posts:       user.Posts,
			Events:      user.Events,
			Friends:     user.Friends,
		},
		"text": "Hello World.",
	})
}

func AuthRouters(route *gin.RouterGroup) {
	auth := route.Group("/auth")
	auth.GET("/refresh", AuthMiddleware.RefreshHandler)
	auth.POST("/login", AuthMiddleware.LoginHandler)
	auth.POST("/register", createUser)
	auth.Use(AuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/logout", AuthMiddleware.LogoutHandler)
	}
}

func AuthUser(username string, password string) (bool, error) {
	user, found, err := database.GetByUsername(username)
	if err != nil {
		return false, err
	} else if !found && err == nil {
		return false, nil
	}
	isCorrect := doPasswordsMatch(user.Credentials.Password, password, []byte(user.Credentials.Hash))
	return isCorrect, nil
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var loginVals Login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	checkPassword, err := AuthUser(userID, password)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	if checkPassword {
		return &database.User{
			Username: userID,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func Authorizer(data interface{}, c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)
	if v, ok := data.(*database.User); ok && v.Username == claims[config.IDENTIFY_KEY].(string) {
		return true
	}
	return false
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*database.User); ok {
		return jwt.MapClaims{
			config.IDENTIFY_KEY: v.Username,
		}
	}
	return jwt.MapClaims{}
}

func IdentifyHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &database.User{
		Username: claims[config.IDENTIFY_KEY].(string),
	}
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
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
