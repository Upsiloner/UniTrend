package user_ctl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Upsiloner/UniTrend/bootstrap"
	"github.com/Upsiloner/UniTrend/domain"
	"github.com/Upsiloner/UniTrend/domain/user_domain"

	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
)

type SignupController struct {
	SignupUsecase user_domain.SignupUsecase
	Env           *bootstrap.Env
	Redis         *redis.Client
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request user_domain.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.NewDefaultResponse(400))
		return
	}

	//判断redis中是否有captcha_suffix对应的项
	_, err = sc.Redis.Get(c.Request.Context(), strings.ToLower(fmt.Sprintf("%s_%s", request.Captcha, request.Suffix))).Result()
	if err != nil {
		c.JSON(http.StatusConflict, domain.NewDefaultResponse(400, "验证码无效"))
		return
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.NewDefaultResponse(400, "该邮箱已被注册"))
		return
	}

	_, err = sc.SignupUsecase.GetUserByName(c, request.Name)
	if err == nil {
		c.JSON(http.StatusConflict, domain.NewDefaultResponse(400, "该用户名已被注册"))
		return
	}

	user := user_domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.NewDefaultResponse(400))
		return
	}

	Token, err := sc.SignupUsecase.CreateJWTToken(&user, sc.Env.TokenSecret, sc.Env.TokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.NewDefaultResponse(400))
		return
	}

	signupResponse := user_domain.SignupResponse{
		DefaultResponse: domain.NewDefaultResponse(200),
		Token:           Token,
		Union_ID:        user.Union_ID,
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshalling user: %v", err)
		c.JSON(http.StatusInternalServerError, domain.NewDefaultResponse(400))
		return
	}

	cmd := sc.Redis.Set(c.Request.Context(), fmt.Sprintf("user_%s", user.Union_ID), userJson, time.Duration(sc.Env.TokenExpiryHour)*time.Hour)
	if cmd.Err() != nil {
		log.Printf("Error setting user in redis: %v", err)
		c.JSON(http.StatusInternalServerError, domain.NewDefaultResponse(400))
		return
	}

	c.JSON(http.StatusOK, signupResponse)
}
