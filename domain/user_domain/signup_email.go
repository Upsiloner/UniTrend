package user_domain

import (
	"context"

	"github.com/Upsiloner/UniTrend/domain"
)

type SignUpEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required"`
}

type SignUpEmailResponse struct {
	domain.DefaultResponse
	Suffix string `json:"suffix"`
}

type SignUpEmailUsecase interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByName(c context.Context, name string) (User, error)
}
