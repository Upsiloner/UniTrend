package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 定义枚举类型表示不同的登录状态

const (
	Success      int = iota // 登录成功, 0
	Unauthorized            // 未授权，即用户名或密码错误, 1
)

func main() {
	r := gin.Default()
	r.POST("/api/login", func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Username == "zhangy" && req.Password == "123456" {
			c.JSON(http.StatusOK, gin.H{"status": Success, "message": "登录成功"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": Unauthorized, "message": "用户名或密码错误"})
		}
	})

	r.Run(":8000")
}
