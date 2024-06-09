package user_uc

import (
	"context"
	"time"

	"github.com/Upsiloner/UniTrend/domain/user_domain"
	"github.com/Upsiloner/UniTrend/internal/token_util"
)

type loginUsecase struct {
	userRepository user_domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository user_domain.UserRepository, timeout time.Duration) user_domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByName(c context.Context, name string) (user_domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetUserByName(ctx, name)
}

func (lu *loginUsecase) CreateJWTToken(user *user_domain.User, secret string, expiry int) (Token string, err error) {
	return token_util.CreateJWTToken(user, secret, expiry)
}
