package service

import (
	"context"
	"donation/entity/client"
)

type UserService interface {
	Create(ctx context.Context, request client.UserCreateRequest) client.UserResponse
	Update(ctx context.Context, request client.UserUpdateRequest) client.UserResponse
	Delete(ctx context.Context, userId int)
	Session(ctx context.Context, request client.UserSessionRequest) client.UserResponse
	FindById(ctx context.Context, userId int) client.UserResponse
	FindByEmail(ctx context.Context, userEmail string) client.UserResponse
	FindAll(ctx context.Context) []client.UserResponse
}
