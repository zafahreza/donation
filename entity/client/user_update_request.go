package client

type UserUpdateRequest struct {
	Id        int    `validate:"required" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `validate:"email" json:"email"`
}
