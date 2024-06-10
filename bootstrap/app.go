package bootstrap

import (
	"log"

	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Application struct {
	Env               *Env
	Postgres          *gorm.DB
	Redis             *redis.Client
	SMTPClientManager *SMTPClientManager
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDatabase(app.Env)
	app.Redis = NewRedisClient(app.Env)
	app.SMTPClientManager = NewSMTPClientManager(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgreSQLConnection(app.Postgres)
}

func (app *Application) CloseRedisConnection() {
	if app.Redis == nil {
		return
	}

	if err := app.Redis.Close(); err != nil {
		log.Fatalf("Failed to close redis connection: %v", err)
	}

	log.Println("Connection to PostgreSQL closed.")
}

func (app *Application) CloseSMTPClientManager() {
	if app.SMTPClientManager == nil {
		return
	}

	if err := app.SMTPClientManager.SMTPClose(); err != nil {
		log.Fatalf("Failed to close SMTP connection: %v", err)
	}
	log.Println("Connection to SMTP server closed.")
}
