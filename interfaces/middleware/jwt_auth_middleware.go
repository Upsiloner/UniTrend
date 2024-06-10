package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Upsiloner/UniTrend/domain"
	"github.com/Upsiloner/UniTrend/internal/token_util"
	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
)

func JwtAuthMiddleware(secret string, rd *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := token_util.IsAuthorized(authToken, secret)
			if authorized {
				UnionID, err := token_util.ExtractUnionIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.NewDefaultResponse(400, fmt.Sprintf("认证Token不合法: %v", err)))
					c.Abort()
					return
				}
				// 查redis判断用户是否已经登出
				userJson, err := rd.Get(c.Request.Context(), fmt.Sprintf("user_%s", UnionID)).Result()
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.NewDefaultResponse(400, "认证已过期"))
					c.Abort()
					return
				}
				c.Set("user", userJson)
				c.Next()
				return
			}
			log.Printf("认证Token无效: %v", err)
			c.JSON(http.StatusUnauthorized, domain.NewDefaultResponse(400, "认证Token无效"))
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.NewDefaultResponse(400, "认证Token无效"))
		c.Abort()
	}
}
