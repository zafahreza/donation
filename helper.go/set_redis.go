package helper

import (
	"context"
	"donation/entity/domain"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"strconv"
)

func SetChacheByUserId(ctx context.Context, chc *redis.Client, user domain.User) {
	if user.Id != 0 {
		fmt.Println("get data by id from db")
		newKey := "userid" + strconv.Itoa(user.Id)
		byteString, err := json.Marshal(user)
		PanicIfError(err)
		err = chc.Set(ctx, newKey, byteString, 0).Err()
		PanicIfError(err)
		fmt.Println("set data by id to redis")
	}
}

func SetChacheByUserEmail(ctx context.Context, chc *redis.Client, user domain.User) {
	if user.Id != 0 {
		fmt.Println("get data by id from db")
		newKey := "userbyemail" + user.Email
		byteString, err := json.Marshal(user)
		PanicIfError(err)
		err = chc.Set(ctx, newKey, byteString, 0).Err()
		PanicIfError(err)
		fmt.Println("set data by email to redis")
	}
}
