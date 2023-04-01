package types

import "github.com/golang-jwt/jwt/v4"

type SignedDetails struct {
	ID string
	jwt.RegisteredClaims
}

type Login struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
