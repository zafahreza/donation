package repository

import (
	"context"
	"donation/entity/domain"
	"donation/helper"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (UserRepository *UserRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, user domain.User, otp domain.OTP) domain.User {
	err := tx.WithContext(ctx).Create(&user).Error
	helper.PanicIfError(err)

	fmt.Println("save new data to db")
	return user
}

func (UserRepository *UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	err := tx.WithContext(ctx).Save(&user).Error
	helper.PanicIfError(err)
	fmt.Println("save update to db")

	return user
}

func (UserRepository *UserRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, user domain.User) {
	err := tx.WithContext(ctx).Delete(&domain.User{}, user.Id).Error
	helper.PanicIfError(err)
	fmt.Println("del data from db")
}

func (UserRepository *UserRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, userId int) (domain.User, error) {
	var user domain.User

	err := tx.WithContext(ctx).Where("id = ?", userId).Find(&user).Error
	helper.PanicIfError(err)

	fmt.Println("get data by id from db")
	return user, nil
}

func (UserRepository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *gorm.DB, userEmail string) (domain.User, error) {
	var user domain.User

	err := tx.WithContext(ctx).Where("email = ?", userEmail).Find(&user).Error
	helper.PanicIfError(err)

	fmt.Println("get data by email from db")
	return user, nil
}

func (UserRepository *UserRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.User {
	users := []domain.User{}
	err := tx.WithContext(ctx).Order("id asc").Find(&users).Order("id desc").Error
	helper.PanicIfError(err)

	fmt.Println("get all data from db")
	return users
}

func (UserRepository *UserRepositoryImpl) UpdateStatusEmail(ctx context.Context, tx *gorm.DB, otp domain.OTP) domain.User {
	var user domain.User

	result := tx.WithContext(ctx).Model(&user).Clauses(clause.Returning{}).Where("email = ?", otp.Email).Update("is_active", true)
	helper.PanicIfError(result.Error)

	fmt.Println("save email active to db")
	return user
}
