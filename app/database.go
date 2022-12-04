package app

import (
	"donation/entity/domain"
	"donation/helper"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewSetupDB() *gorm.DB {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)
	host := os.Getenv("HOST")
	username := os.Getenv("NAME")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("PASS_WORD")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, username, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	pool, err := db.DB()
	helper.PanicIfError(err)

	pool.SetMaxIdleConns(5)
	pool.SetMaxOpenConns(20)
	pool.SetConnMaxIdleTime(10 * time.Minute)
	pool.SetConnMaxLifetime(60 * time.Minute)

	db.AutoMigrate(&domain.User{})

	return db
}
