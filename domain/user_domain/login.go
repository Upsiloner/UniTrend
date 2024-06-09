package user_domain

import (
	"context"

	"github.com/Upsiloner/UniTrend/domain"
)

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	domain.DefaultResponse
	Token    string `json:"token"`
	Union_ID string `json:"union_id"`
}

type LoginUsecase interface {
	GetUserByName(c context.Context, name string) (User, error)
	CreateJWTToken(user *User, secret string, expiry int) (Token string, err error)
}
