package auth

type LoginDto struct {
	email    string `validate:"required,email"`
	password string
}

type TokenType struct {
	accessToken  string `json:access_token`
	refreshToken string `json:refresh_token`
}

type Token struct {
	TokenString string `validate:"required"`
	UserId      uint   `validate:"required"`
	Expires     int64  `validate:"required"`
	Blacklisted bool   `validate:"required"`
}
