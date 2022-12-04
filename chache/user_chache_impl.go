package chache

import (
	"context"
	"donation/entity/domain"
	"donation/helper"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

type UserChacheImpl struct {
	Redis *redis.Client
}

func NewUserChache(chache *redis.Client) UserChache {
	return &UserChacheImpl{Redis: chache}
}

func (c UserChacheImpl) Set(ctx context.Context, user domain.User, key string) {
	userMarshal := helper.Marshal(user)
	err := c.Redis.Set(ctx, key, userMarshal, 0).Err()
	helper.PanicIfError(err)
	fmt.Println("set chache to redis")
}

func (c UserChacheImpl) Get(ctx context.Context, key string) (domain.User, error) {
	var user domain.User

	result, err := c.Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return user, err
	}
	helper.PanicIfError(err)

	user = helper.UnMarshalUser(result)
	fmt.Println("get chache from redis")
	return user, nil
}

func (c UserChacheImpl) Del(ctx context.Context, keys ...string) {
	args := make([]string, len(keys))
	for i, key := range keys {
		args[0+i] = key
	}
	err := c.Redis.Del(ctx, args...).Err()
	helper.PanicIfError(err)
	fmt.Println("del chache from redis")
}

func (c UserChacheImpl) SetOtp(ctx context.Context, otp domain.OTP) {
	key := "otpfor" + otp.Email

	userMarshal := helper.Marshal(otp)
	err := c.Redis.Set(ctx, key, userMarshal, 60*time.Second).Err()
	helper.PanicIfError(err)
	fmt.Println("set otp to redis")
}

func (c UserChacheImpl) GetOtp(ctx context.Context, otp domain.OTP) (domain.OTP, error) {
	key := "otpfor" + otp.Email

	result, err := c.Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return otp, err
	}
	helper.PanicIfError(err)

	otp = helper.UnMarshalOtp(result)
	fmt.Println("get otp from redis")
	return otp, nil
}
