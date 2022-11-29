package service

import (
	"donation/entity/client"
)

type UserService interface {
	Create(request client.UserCreateRequest) client.UserResponse
	Update(request client.UserUpdateRequest) client.UserResponse
	Delete(userId int)
	Session(request client.UserSessionRequest) client.UserResponse
	FindById(userId int) client.UserResponse
	FindByEmail(userEmail string) client.UserResponse
	FindAll() []client.UserResponse
}
