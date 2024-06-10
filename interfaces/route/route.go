package route

import (
	"time"

	"github.com/Upsiloner/UniTrend/bootstrap"
	"github.com/Upsiloner/UniTrend/interfaces/middleware"
	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, rd *redis.Client, SMTPClientManager *bootstrap.SMTPClientManager, gin *gin.Engine) {
	publicUserRouter := gin.Group("/api/user")
	// All User Public APIs
	NewLoginRouter(env, timeout, db, rd, publicUserRouter)
	NewSignupRouter(env, timeout, db, rd, publicUserRouter)
	NewSignUpEmailRouter(env, timeout, db, rd, SMTPClientManager, publicUserRouter)

	protectedUserRouter := gin.Group("/api/user")
	// Middleware to verify AccessToken
	protectedUserRouter.Use(middleware.JwtAuthMiddleware(env.TokenSecret, rd))
	// All User Private APIs
	NewGetUserRouter(timeout, db, protectedUserRouter)
}
