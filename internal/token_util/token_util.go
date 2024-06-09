package token_util

import (
	"fmt"
	"time"

	"github.com/Upsiloner/UniTrend/domain"
	"github.com/Upsiloner/UniTrend/domain/user_domain"
	jwt "github.com/golang-jwt/jwt/v4"
)

// 创建一个包含常用用户信息的JWT Token
func CreateJWTToken(user *user_domain.User, secret string, expiry int) (accessToken string, err error) {
	exp := jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry)))
	claims := &domain.JwtCustomClaims{
		Name:     user.Name,
		Union_ID: user.Union_ID,
		Status:   fmt.Sprint(user.Status),
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractUnionIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}

	return claims["union_id"].(string), nil
}
