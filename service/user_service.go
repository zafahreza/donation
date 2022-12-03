package service

import (
	"context"
	"donation/entity/client"
)

type UserService interface {
	Create(ctx context.Context, request client.UserCreateRequest) client.UserResponse
	Update(ctx context.Context, request client.UserUpdateRequest) client.UserResponse
	Delete(ctx context.Context, userId int)
	Session(ctx context.Context, request client.UserSessionRequest) client.UserLoginResponse
	FindById(ctx context.Context, userId int) client.UserResponse
	FindByEmail(ctx context.Context, userEmail string) client.UserResponse
	FindAll(ctx context.Context) []client.UserResponse
	FindOtp(ctx context.Context, request client.UserOtpRequest) client.UserResponse
	GetNewOtp(ctx context.Context, request client.UserGetNewOtpRequest) client.UserGetNewOtpResponse
}
