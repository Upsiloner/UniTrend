package user_uc

import (
	"context"
	"time"

	"github.com/Upsiloner/UniTrend/domain/user_domain"
)

type SignUpEmailUsecase struct {
	userRepository user_domain.UserRepository
	contextTimeout time.Duration
}

func NewSignUpEmailUsecase(userRepository user_domain.UserRepository, timeout time.Duration) user_domain.SignUpEmailUsecase {
	return &SignUpEmailUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *SignUpEmailUsecase) GetUserByEmail(c context.Context, email string) (user_domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetUserByEmail(ctx, email)
}

func (su *SignUpEmailUsecase) GetUserByName(c context.Context, name string) (user_domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetUserByName(ctx, name)
}
