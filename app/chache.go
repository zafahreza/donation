package app

import (
	"donation/helper.go"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
	"os"
)

func NewChacheDB() *redis.Client {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)
	host := os.Getenv("REDIS_HOST")
	username := os.Getenv("REDIS_USER")
	db := os.Getenv("REDIS_DB")
	password := os.Getenv("REDIS_PASS")
	port := os.Getenv("REDIS_PORT")

	redisUrl := fmt.Sprintf("redis://%s:%s@%s:%s/%s", username, password, host, port, db)

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic("gagal connect")
	}

	rdb := redis.NewClient(opt)

	return rdb
}
