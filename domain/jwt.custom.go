package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Name     string `json:"name"`
	Union_ID string `json:"union_id"`
	Status   string `json:"status"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
