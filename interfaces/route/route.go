package route

import (
	"time"

	"github.com/Upsiloner/UniTrend/bootstrap"
	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, rd *redis.Client, gin *gin.Engine) {
	publicUserRouter := gin.Group("/api/user")
	// All User Public APIs
	NewLoginRouter(env, timeout, db, rd, publicUserRouter)

}
