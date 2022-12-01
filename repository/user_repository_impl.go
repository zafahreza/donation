package repository

import (
	"donation/entity/domain"
	"donation/helper.go"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (UserRepository *UserRepositoryImpl) Save(tx *gorm.DB, user domain.User) domain.User {
	err := tx.Create(&user).Error
	helper.PanicIfError(err)

	return user

}

func (UserRepository *UserRepositoryImpl) Update(tx *gorm.DB, user domain.User) domain.User {
	err := tx.Save(&user).Error
	helper.PanicIfError(err)

	return user
}

func (UserRepository *UserRepositoryImpl) Delete(tx *gorm.DB, user domain.User) {
	err := tx.Delete(&domain.User{}, user.Id).Error
	helper.PanicIfError(err)
}

func (UserRepository *UserRepositoryImpl) FindById(tx *gorm.DB, userId int) (domain.User, error) {
	user := domain.User{}

	err := tx.Where("id = ?", userId).Find(&user).Error
	helper.PanicIfError(err)

	return user, nil
}

func (UserRepository *UserRepositoryImpl) FindByEmail(tx *gorm.DB, userEmail string) (domain.User, error) {
	user := domain.User{}
	err := tx.Where("email = ?", userEmail).Find(&user).Error
	helper.PanicIfError(err)

	return user, nil
}

func (UserRepository *UserRepositoryImpl) FindAll(tx *gorm.DB) []domain.User {
	users := []domain.User{}
	err := tx.Order("id asc").Find(&users).Order("id desc").Error
	helper.PanicIfError(err)

	return users
}
