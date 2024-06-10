package user_domain

import (
	"context"

	"github.com/Upsiloner/UniTrend/domain"
)

type SignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Captcha  string `json:"captcha" binding:"required"`
	Suffix   string `json:"suffix" binding:"required"`
}

type SignupResponse struct {
	domain.DefaultResponse
	Token    string `json:"token"`
	Union_ID string `json:"union_id"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByName(c context.Context, name string) (User, error)
	CreateJWTToken(user *User, secret string, expiry int) (Token string, err error)
}
