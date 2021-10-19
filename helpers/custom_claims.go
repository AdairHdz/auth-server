package helpers

import (
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	UserID string
	UserType int
	EmailAddress string
	jwt.RegisteredClaims
}