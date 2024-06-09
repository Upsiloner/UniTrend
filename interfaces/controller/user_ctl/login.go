package user_ctl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Upsiloner/UniTrend/bootstrap"
	"github.com/Upsiloner/UniTrend/domain"
	"github.com/Upsiloner/UniTrend/domain/user_domain"

	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
)

type LoginController struct {
	LoginUsecase user_domain.LoginUsecase
	Env          *bootstrap.Env
	Redis        *redis.Client
}

func (lc *LoginController) Login(c *gin.Context) {
	var request user_domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.NewDefaultResponse(400))
		return
	}

	user, err := lc.LoginUsecase.GetUserByName(c, request.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.NewDefaultResponse(400, "不存在该用户"))
		return
	}

	if user.Password != request.Password {
		c.JSON(http.StatusUnauthorized, domain.NewDefaultResponse(400, "用户名或密码错误"))
		return
	}

	Token, err := lc.LoginUsecase.CreateJWTToken(&user, lc.Env.TokenSecret, lc.Env.TokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.NewDefaultResponse(400))
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error marshalling user: %v", err)
		return
	}
	cmd := lc.Redis.Set(c.Request.Context(), fmt.Sprintf("user_%s", user.Union_ID), userJson, time.Duration(lc.Env.TokenExpiryHour)*time.Hour)
	if cmd.Err() != nil {
		log.Printf("Error setting user in redis: %v", cmd.Err())
		c.JSON(http.StatusInternalServerError, domain.NewDefaultResponse(400))
		return
	}

	loginResponse := user_domain.LoginResponse{
		DefaultResponse: domain.NewDefaultResponse(200),
		Token:           Token,
		Union_ID:        user.Union_ID,
	}

	c.JSON(http.StatusOK, loginResponse)
}
