package repository

import (
	"context"
	"donation/entity/domain"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, chache *redis.Client, tx *gorm.DB, user domain.User, otp domain.OTP) domain.User
	Update(ctx context.Context, chache *redis.Client, tx *gorm.DB, user domain.User) domain.User
	Delete(ctx context.Context, chache *redis.Client, tx *gorm.DB, user domain.User)
	FindById(ctx context.Context, chache *redis.Client, tx *gorm.DB, userId int) (domain.User, error)
	FindByEmail(ctx context.Context, chache *redis.Client, tx *gorm.DB, UserEmail string) (domain.User, error)
	FindAll(ctx context.Context, tx *gorm.DB) []domain.User
	FindOTp(ctx context.Context, chache *redis.Client, otp domain.OTP) (domain.OTP, error)
	SetOTp(ctx context.Context, chache *redis.Client, otp domain.OTP)
	DelOTP(ctx context.Context, chache *redis.Client, otp domain.OTP)
	UpdateStatusEmail(ctx context.Context, tx *gorm.DB, otp domain.OTP) domain.User
}
