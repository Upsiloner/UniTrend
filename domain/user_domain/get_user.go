package user_domain

import (
	"context"

	"github.com/Upsiloner/UniTrend/domain"
)

type GetUserRequest struct {
	Union_ID string `uri:"union_id" binding:"required"`
}

type GetUserResponse struct {
	domain.DefaultResponse
	User User `json:"user"`
}

type GetUserUsecase interface {
	GetUserByUnionID(c context.Context, union_id string) (User, error)
}
