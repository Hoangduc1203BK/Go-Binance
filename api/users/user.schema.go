package users



type createUserResponse struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

type updateUserRequest struct {
	createUserRequest
}

type GetUserDto struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
