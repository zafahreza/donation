package repository

import (
	"donation/entity/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(tx *gorm.DB, user domain.User) domain.User
	Update(tx *gorm.DB, user domain.User) domain.User
	Delete(tx *gorm.DB, user domain.User)
	FindById(tx *gorm.DB, userId int) (domain.User, error)
	FindByEmail(tx *gorm.DB, UserEmail string) (domain.User, error)
	FindAll(tx *gorm.DB) []domain.User
}
