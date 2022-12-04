package helper

import (
	"donation/entity/client"
	"donation/entity/domain"
)

func ToUserResponse(user domain.User) client.UserResponse {
	return client.UserResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Bio:       user.Bio,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []domain.User) []client.UserResponse {
	var userResponses []client.UserResponse

	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToUserLoginResponse(user domain.User, token string) client.UserLoginResponse {
	return client.UserLoginResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Bio:       user.Bio,
		Token:     token,
		UpdatedAt: user.UpdatedAt,
	}
}
