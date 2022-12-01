package client

type UserCreateRequest struct {
	FirstName string `validate:"required" json:"first_name"`
	LastName  string `validate:"required" json:"last_name"`
	Password  string `validate:"required" json:"password"`
	Email     string `validate:"required,email" json:"email"`
}
