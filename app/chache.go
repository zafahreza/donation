package app

import (
	"donation/helper"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

func NewChacheDB() *redis.Client {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)
	host := os.Getenv("REDIS_HOST")
	username := os.Getenv("REDIS_USER")
	db := os.Getenv("REDIS_DB")
	password := os.Getenv("REDIS_PASS")
	port := os.Getenv("REDIS_PORT")

	dbInt, err := strconv.Atoi(db)
	helper.PanicIfError(err)

	//redisUrl := fmt.Sprintf("redis://%s:%s@%s:%s/%s", username, password, host, port, db)
	//
	//opt, err := redis.ParseURL(redisUrl)
	//if err != nil {
	//	panic("gagal connect")
	//}

	rdb := redis.NewClient(&redis.Options{
		Addr:            host + ":" + port,
		Password:        password,
		Username:        username,
		DB:              dbInt,
		MinIdleConns:    5,
		ConnMaxIdleTime: 10 * time.Minute,
		ConnMaxLifetime: 60 * time.Minute,
		PoolSize:        20,
	})

	return rdb
}
