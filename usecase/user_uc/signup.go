package user_uc

import (
	"context"
	"time"

	"github.com/Upsiloner/UniTrend/domain/user_domain"
	"github.com/Upsiloner/UniTrend/internal/token_util"
)

type signupUsecase struct {
	userRepository user_domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository user_domain.UserRepository, timeout time.Duration) user_domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *user_domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (user_domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetUserByEmail(ctx, email)
}

func (su *signupUsecase) GetUserByName(c context.Context, name string) (user_domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetUserByName(ctx, name)
}

func (su *signupUsecase) CreateJWTToken(user *user_domain.User, secret string, expiry int) (accessToken string, err error) {
	return token_util.CreateJWTToken(user, secret, expiry)
}
