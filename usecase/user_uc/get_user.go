package user_uc

import (
	"context"
	"time"

	"github.com/Upsiloner/UniTrend/domain/user_domain"
)

type getUserUsecase struct {
	userRepository user_domain.UserRepository
	contextTimeout time.Duration
}

func NewGetUserUsecase(userRepository user_domain.UserRepository, timeout time.Duration) user_domain.GetUserUsecase {
	return &getUserUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (guu *getUserUsecase) GetUserByUnionID(c context.Context, union_id string) (user_domain.User, error) {
	ctx, cancel := context.WithTimeout(c, guu.contextTimeout)
	defer cancel()
	return guu.userRepository.GetUserByUnionID(ctx, union_id)
}
