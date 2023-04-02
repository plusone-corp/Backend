package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"plusone/backend/config"
	"plusone/backend/database"
	"plusone/backend/errorHandler"
	"plusone/backend/types"
	"strings"
	"time"
)

var (
	RefreshKey = []byte(config.RF_JWT_SECRET)
	AccessKey  = []byte(config.JWT_SECRET)
)

func Sign(userId primitive.ObjectID) (*Tokens, error) {

	acClaim := &types.SignedDetails{
		ID: userId.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "PlusOne",
		},
	}

	rfClaim := &types.SignedDetails{
		ID: userId.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "PlusOne",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, acClaim)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rfClaim)

	actString, err := accessToken.SignedString(AccessKey)
	if err != nil {
		return nil, err
	}
	rfString, err := refreshToken.SignedString(RefreshKey)
	if err != nil {
		return nil, err
	}

	tokens := Tokens{
		AccessToken:  actString,
		RefreshToken: rfString,
	}

	err = database.UpdateRefreshToken(userId, rfString)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}

func ParseAccessToken(tokenStr string) (*types.SignedDetails, bool, *string) {
	token, err := jwt.ParseWithClaims(tokenStr, &types.SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return AccessKey, nil
	})
	if err != nil {
		return nil, false, &errorHandler.FailedTokenValidation
	}

	claims := token.Claims.(*types.SignedDetails)

	return claims, token.Valid, nil
}

func ParseRefreshToken(tokenStr string) (*types.SignedDetails, bool, *string) {
	token, err := jwt.ParseWithClaims(tokenStr, &types.SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return RefreshKey, nil
	})
	if err != nil {
		return nil, false, &errorHandler.FailedTokenValidation
	}

	claims := token.Claims.(*types.SignedDetails)

	return claims, token.Valid, nil
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := validateHeaders(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": errorHandler.AuthorizationKeyNotFound,
			})
			return
		}

		claims, valid, parseErr := ParseAccessToken(*token)
		if parseErr != nil && claims == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": errorHandler.InvalidToken,
			})
			return
		} else if !valid && parseErr != nil {
			if claims.ExpiresAt.Unix() > time.Now().Unix() {
				c.AbortWithStatusJSON(http.StatusRequestTimeout, gin.H{
					"status":  http.StatusRequestTimeout,
					"message": errorHandler.RefreshToken,
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusForbidden,
				"message": *parseErr,
			})
			return
		}

		savedToken, found := LoggedOutToken.Get(claims.ID)
		if found && savedToken == *token {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": errorHandler.InvalidToken,
			})
			return
		}

		c.Set("JWT_PAYLOAD", claims)

		c.Next()
	}
}

func validateHeaders(header string) (*string, *string) {
	log.Println(header)
	if !strings.HasPrefix(header, "Bearer") {
		return nil, &errorHandler.InvalidMethod
	}

	tokenString := strings.Split(header, " ")[1]
	return &tokenString, nil
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
