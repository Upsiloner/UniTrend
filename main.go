package main

import (
	"time"

	"github.com/Upsiloner/UniTrend/bootstrap"
	route "github.com/Upsiloner/UniTrend/interfaces/route"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Postgres
	redis := app.Redis
	SMTPClientManager := app.SMTPClientManager

	// Deconstruction
	defer app.CloseDBConnection()
	defer app.CloseRedisConnection()
	defer app.CloseSMTPClientManager()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, redis, SMTPClientManager, gin)

	gin.Run(env.ServerAddress)
}
