package client

type UserUpdateRequest struct {
	Id        int    `validate:"required" json:"id"`
	FirstName string `validate:"required" json:"first_name"`
	LastName  string `validate:"required" json:"last_name"`
	Email     string `validate:"required,email" json:"email"`
	Bio       string `json:"bio"`
}
