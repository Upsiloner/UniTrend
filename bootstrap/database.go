package bootstrap

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(env *Env) *gorm.DB {
	// 构建DSN字符串
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		env.DBHost, env.DBUser, env.DBPass, env.DBName, env.DBPort)

	// 使用Gorm开启PostgreSQL连接
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移数据库模式以确保模式是最新的
	if err := db.AutoMigrate(&user_domain.User{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	// 尝试连接，确保数据库是可达的
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
