package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"plusone/backend/auth"
	"plusone/backend/config"
	"plusone/backend/user"
	"time"
)

var (
	Router *gin.Engine
)

func main() {
	// Init
	Router = gin.Default()
	Router.GET("/", hello)
	Router.NoRoute(handleNotFound)

	var err error
	auth.AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "PlusOne",
		Key:             []byte(config.JWT_SECRET),
		Timeout:         time.Minute * time.Duration(config.JWT_TIMEOUT_TIME),
		MaxRefresh:      time.Hour * time.Duration(config.JWT_REFRESH_TIME),
		IdentityKey:     config.IDENTIFY_KEY,
		PayloadFunc:     auth.PayloadFunc,
		IdentityHandler: auth.IdentifyHandler,
		Authenticator:   auth.Authenticator,
		Authorizator:    auth.Authorizer,
		Unauthorized:    auth.Unauthorized,
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := auth.AuthMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("AuthMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	apiGroup := Router.Group("/api")
	auth.AuthRouters(apiGroup)
	user.UserHandler(apiGroup)

	err = Router.Run()
	if err != nil {
		panic(err)
	}
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}

func handleNotFound(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{
		"status":  404,
		"message": "Page not found!",
	})
}
