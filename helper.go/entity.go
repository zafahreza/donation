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
