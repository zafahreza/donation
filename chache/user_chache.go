package chache

import (
	"context"
	"donation/entity/domain"
)

type UserChache interface {
	Set(ctx context.Context, user domain.User, key string)
	Get(ctx context.Context, key string) (domain.User, error)
	Del(ctx context.Context, keys ...string)
	SetOtp(ctx context.Context, otp domain.OTP)
	GetOtp(ctx context.Context, otp domain.OTP) (domain.OTP, error)
}
