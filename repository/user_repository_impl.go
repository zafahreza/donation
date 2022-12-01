package repository

import (
	"context"
	"donation/entity/domain"
	"donation/helper.go"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (UserRepository *UserRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	err := tx.WithContext(ctx).Create(&user).Error
	helper.PanicIfError(err)

	return user

}

func (UserRepository *UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	err := tx.WithContext(ctx).Save(&user).Error
	helper.PanicIfError(err)

	return user
}

func (UserRepository *UserRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, user domain.User) {
	err := tx.WithContext(ctx).Delete(&domain.User{}, user.Id).Error
	helper.PanicIfError(err)
}

func (UserRepository *UserRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, userId int) (domain.User, error) {
	user := domain.User{}

	err := tx.WithContext(ctx).Where("id = ?", userId).Find(&user).Error
	helper.PanicIfError(err)

	return user, nil
}

func (UserRepository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *gorm.DB, userEmail string) (domain.User, error) {
	user := domain.User{}
	err := tx.WithContext(ctx).Where("email = ?", userEmail).Find(&user).Error
	helper.PanicIfError(err)

	return user, nil
}

func (UserRepository *UserRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.User {
	users := []domain.User{}
	err := tx.WithContext(ctx).Order("id asc").Find(&users).Order("id desc").Error
	helper.PanicIfError(err)

	return users
}
