package users


type createUserResponse struct {
	Id 			int    `json: Id`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

type updateUserRequest struct {
	Name        string `json:"name" ,alphanum"`
	Password    string `json:"password" ,min=6"`
	PhoneNumber string `json:"phone_number" ,number,min=10"`
	Email       string `json:"email" ,email"`
}

type GetUserDto struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type createUserRequest struct {
	Name        string `json:"name" binding:"required,alphanum"`
	Password    string `json:"password" binding:"required,min=6"`
	PhoneNumber string `json:"phone_number" binding:"required,number,min=10"`
	Email       string `json:"email" binding:"required,email"`
}

// type updateUserRequest struct {
	
// }
