package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"plusone/backend/config"
	"plusone/backend/database"
	"plusone/backend/types"
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
	user, found, error := database.GetByUsername(username.(*types.User).Username)
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
		"user": types.User{
			Username:    user.Username,
			Email:       user.Email,
			Age:         user.Age,
			Location:    user.Location,
			DisplayName: user.DisplayName,
			CreatedAt:   user.CreatedAt,
			Avatar:      user.Avatar,
			Posts:       user.Posts,
			Events:      user.Events,
			Friends:     user.Friends,
			Description: user.Description,
			Level:       user.Level,
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
