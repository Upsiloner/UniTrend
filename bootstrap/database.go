package bootstrap

import (
	"fmt"
	"log"

	"UniTrend/domain/user_domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(env *Env) *gorm.DB {
	// Building DSN strings
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		env.DBHost, env.DBUser, env.DBPass, env.DBName, env.DBPort)

	// Enable PostgreSQL connection using Gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Automatically migrate database schema to ensure it is up-to-date
	if err := db.AutoMigrate(&user_domain.User{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	// Attempt to connect and ensure that the database is reachable
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to PostgreSQL database.")
	return db
}

func ClosePostgreSQLConnection(db *gorm.DB) {
	if db == nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection for closing: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}

	log.Println("Connection to PostgreSQL closed.")
}
