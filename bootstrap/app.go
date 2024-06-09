package bootstrap

import (
	"log"

	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Application struct {
	Env      *Env
	Postgres *gorm.DB
	Redis    *redis.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDatabase(app.Env)
	app.Redis = NewRedisClient(app.Env)
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
