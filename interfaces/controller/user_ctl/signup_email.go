package user_ctl

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Upsiloner/UniTrend/bootstrap"
	"github.com/Upsiloner/UniTrend/domain"
	"github.com/Upsiloner/UniTrend/domain/user_domain"
	"github.com/Upsiloner/UniTrend/internal/util"
	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
	mail "github.com/xhit/go-simple-mail/v2"
)

type SignUpEmailController struct {
	SignUpEmailUsecase user_domain.SignUpEmailUsecase
	Timeout            time.Duration
	Env                *bootstrap.Env
	Redis              *redis.Client
	SMTPClientManager  *bootstrap.SMTPClientManager
}

func (sec *SignUpEmailController) SignUpEmail(c *gin.Context) {
	var request user_domain.SignUpEmailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding request: %v", err)
		c.JSON(http.StatusBadRequest, domain.NewDefaultResponse(400))
		return
	}

	// 创建一个带有超时的上下文
	ctx, cancel := context.WithTimeout(c, sec.Timeout)
	defer cancel()

	_, err := sec.SignUpEmailUsecase.GetUserByEmail(ctx, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.NewDefaultResponse(400, "该邮箱已被注册"))
		return
	}

	_, err = sec.SignUpEmailUsecase.GetUserByName(ctx, request.Name)
	if err == nil {
		c.JSON(http.StatusConflict, domain.NewDefaultResponse(400, "该用户名已被注册"))
		return
	}

	captcha, err := util.GenerateUUID(6)
	if err != nil {
		log.Printf("Error generating captcha: %v", err)
		c.JSON(http.StatusInternalServerError, domain.NewDefaultResponse(500))
		return
	}
	suffix, err := util.GenerateUUID(4)
	if err != nil {
		log.Printf("Error generating suffix: %v", err)
		c.JSON(http.StatusInternalServerError, domain.NewDefaultResponse(500))
		return
	}

	email := mail.NewMSG()
	email.SetFrom("UniTrace <zhangyutongxue@163.com>").
		AddTo(request.Email).
		SetSubject("[UniTrace Sign Up] 查看你的邮箱注册验证码").
		SetBody(mail.TextHTML, getSignUpCaptchaHTML(captcha))

	// 获取SMTP客户端并发送邮件
	smtpClient := sec.SMTPClientManager.GetClient()
	if err := email.Send(smtpClient); err != nil {
		log.Printf("Error sending email: %v", err)
		c.JSON(http.StatusInternalServerError, domain.NewDefaultResponse(500, "发送邮件失败"))
		return
	}

	if err := sec.Redis.Set(ctx, strings.ToLower(fmt.Sprintf("%s_%s", captcha, suffix)), request.Email, time.Duration(sec.Env.CaptchaExpiryMinute)*time.Minute).Err(); err != nil {
		log.Printf("Error setting email captcha in redis: %v", err)
		c.JSON(http.StatusInternalServerError, domain.NewDefaultResponse(500))
		return
	}

	SignUpEmailResponse := user_domain.SignUpEmailResponse{
		DefaultResponse: domain.NewDefaultResponse(200),
		Suffix:          suffix,
	}

	c.JSON(http.StatusOK, SignUpEmailResponse)
}

func getSignUpCaptchaHTML(captcha string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
		<html lang="zh-CN">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>注册验证码</title>
			<style>
				html,
				body {
					margin: 16px;
					padding: 8px;
					font-family: Arial, sans-serif;
					background-color: #f4f4f9;
					color: #333;
					line-height: 1.6;
				}
				.container {
					max-width: 600px;
					margin: 30px auto;
					padding: 20px;
					background-color: #f9f9f9;
					box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
					border-radius: 10px;
				}

				.header {
					text-align: center;
					padding: 10px 0;
					border-bottom: 2px solid rgb(186, 164, 201);
				}

				.header h1 {
					margin: 0;
					color: #333;
				}

				.content {
					margin-top: 20px;
					padding: 20px;
					text-align: center;
				}

				.content h2 {
					font-weight: bold;
					background-color: #e8e8e8;
					font-size: 26px;
					color: purple;
				}

				.footer {
					text-align: center;
					padding: 10px 0;
					border-top: 1px solid #eee;
					font-size: 12px;
					color: #aaa;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>邮箱验证码</h1>
				</div>
				<div class="content">
					<p>尊敬的用户，您正在进行邮箱验证操作</p>
					<p>您的注册验证码是：</p>
					<h2>%s</h2>
					<p>此验证码在 3 分钟内有效</p>
					<p>请保存好您的信息，谨防诈骗！</p>
					<p>感谢您的使用！</p>
				</div>
				<div class="footer">
					<p>&copy; 2024 UniTrace. 保留所有权利</p>
				</div>
			</div>
		</body>
		</html>`, captcha)
}
