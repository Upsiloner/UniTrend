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
	group.POST("/api/login", lc.Login)
}
