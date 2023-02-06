package utils

import (
	"github.com/golang-jwt/jwt/v4"
)

var (
	EmqxKey    = "1f9c5b734fe27865"
	EmqxSecret = "lV9C2iefOp9Cr9BeiB5rr3N9CBolJjKk3HruhqEpHQxsuVD"
)

type M map[string]interface{}

type UserClaim struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var (
	JwtKey = "my-jwt-secret-key"
)
