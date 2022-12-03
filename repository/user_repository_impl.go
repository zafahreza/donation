package repository

import (
	"context"
	"donation/entity/domain"
	"donation/helper.go"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (UserRepository *UserRepositoryImpl) Save(ctx context.Context, chache *redis.Client, tx *gorm.DB, user domain.User, otp domain.OTP) domain.User {
	err := tx.WithContext(ctx).Create(&user).Error
	helper.PanicIfError(err)

	otpMarshal, err := json.Marshal(otp)

	key := "otpfor" + otp.Email

	err = chache.Set(ctx, key, otpMarshal, 60*time.Second).Err()
	helper.PanicIfError(err)

	fmt.Println("save new data to db")

	return user

}

func (UserRepository *UserRepositoryImpl) Update(ctx context.Context, chache *redis.Client, tx *gorm.DB, user domain.User) domain.User {
	err := tx.WithContext(ctx).Save(&user).Error
	helper.PanicIfError(err)
	fmt.Println("save update to db")

	helper.SetChacheByUserId(ctx, chache, user)
	helper.SetChacheByUserEmail(ctx, chache, user)
	return user
}

func (UserRepository *UserRepositoryImpl) Delete(ctx context.Context, chache *redis.Client, tx *gorm.DB, user domain.User) {
	err := tx.WithContext(ctx).Delete(&domain.User{}, user.Id).Error
	helper.PanicIfError(err)
	fmt.Println("del data from db")

	key := "userid" + strconv.Itoa(user.Id)
	chache.Del(ctx, key)
	fmt.Println("del data by id from redis")

	key2 := "userbyemail" + user.Email
	chache.Del(ctx, key2)
	fmt.Println("del data by email from redis")
}

func (UserRepository *UserRepositoryImpl) FindById(ctx context.Context, chache *redis.Client, tx *gorm.DB, userId int) (domain.User, error) {
	var user domain.User

	key := "userid" + strconv.Itoa(userId)
	result, err := chache.Get(ctx, key).Result()
	if err == redis.Nil {
		err := tx.WithContext(ctx).Where("id = ?", userId).Find(&user).Error
		helper.PanicIfError(err)

		helper.SetChacheByUserId(ctx, chache, user)

		return user, nil
	}

	user = helper.UnMarshal(result)
	fmt.Println("get data by id from redis")
	return user, nil
}

func (UserRepository *UserRepositoryImpl) FindByEmail(ctx context.Context, chache *redis.Client, tx *gorm.DB, userEmail string) (domain.User, error) {
	var user domain.User

	key := "userbyemail" + userEmail
	result, err := chache.Get(ctx, key).Result()
	if err == redis.Nil {
		err := tx.WithContext(ctx).Where("email = ?", userEmail).Find(&user).Error
		helper.PanicIfError(err)

		helper.SetChacheByUserEmail(ctx, chache, user)

		return user, nil
	}

	user = helper.UnMarshal(result)
	fmt.Println("get data by email from redis")
	return user, nil
}

func (UserRepository *UserRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.User {
	users := []domain.User{}
	err := tx.WithContext(ctx).Order("id asc").Find(&users).Order("id desc").Error
	helper.PanicIfError(err)

	return users
}

func (UserRepository *UserRepositoryImpl) FindOTp(ctx context.Context, chache *redis.Client, otp domain.OTP) (domain.OTP, error) {
	key := "otpfor" + otp.Email
	result, err := chache.Get(ctx, key).Result()
	if err == redis.Nil {
		return otp, err
	}
	err = json.Unmarshal([]byte(result), &otp)
	helper.PanicIfError(err)

	return otp, nil

}

func (UserRepository *UserRepositoryImpl) DelOTP(ctx context.Context, chache *redis.Client, otp domain.OTP) {

	key := "otpfor" + otp.Email

	chache.Del(ctx, key)
	fmt.Println("del data otp from redis")

}

func (UserRepository *UserRepositoryImpl) UpdateStatusEmail(ctx context.Context, tx *gorm.DB, otp domain.OTP) domain.User {
	var user domain.User

	result := tx.WithContext(ctx).Model(&user).Clauses(clause.Returning{}).Where("email = ?", otp.Email).Update("is_active", true)
	helper.PanicIfError(result.Error)

	return user
}

func (UserRepository *UserRepositoryImpl) SetOTp(ctx context.Context, chache *redis.Client, otp domain.OTP) {
	otpMarshal, err := json.Marshal(otp)

	key := "otpfor" + otp.Email

	err = chache.Set(ctx, key, otpMarshal, 60*time.Second).Err()

	key2 := "userbyemail" + otp.Email
	chache.Del(ctx, key2)
	helper.PanicIfError(err)
}
