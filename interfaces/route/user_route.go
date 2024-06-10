package route

import (
	"time"

	"github.com/Upsiloner/UniTrend/bootstrap"
	"github.com/Upsiloner/UniTrend/interfaces/controller/user_ctl"
	"github.com/Upsiloner/UniTrend/repository"
	"github.com/Upsiloner/UniTrend/usecase/user_uc"
	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, rd *redis.Client, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := user_ctl.LoginController{
		LoginUsecase: user_uc.NewLoginUsecase(ur, timeout),
		Env:          env,
		Redis:        rd,
	}
	group.POST("/login", lc.Login)
}

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, rd *redis.Client, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sc := user_ctl.SignupController{
		SignupUsecase: user_uc.NewSignupUsecase(ur, timeout),
		Env:           env,
		Redis:         rd,
	}
	group.POST("/signup", sc.Signup)
}

func NewSignUpEmailRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, rd *redis.Client, SMTPClientManager *bootstrap.SMTPClientManager, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sec := user_ctl.SignUpEmailController{
		SignUpEmailUsecase: user_uc.NewSignUpEmailUsecase(ur, timeout),
		Timeout:            timeout,
		Env:                env,
		Redis:              rd,
		SMTPClientManager:  SMTPClientManager,
	}
	group.POST("/signupemail", sec.SignUpEmail)
}

func NewGetUserRouter(timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	guc := user_ctl.GetUserController{
		GetUserUsecase: user_uc.NewGetUserUsecase(ur, timeout),
	}
	group.GET("/:union_id", guc.GetUser)
}
