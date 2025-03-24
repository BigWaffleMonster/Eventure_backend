package http_models

type UserRegisterInput struct {
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type UserLoginInput struct {
	Email    string  `json:"email"`
	Password *string `json:"password"`
}
