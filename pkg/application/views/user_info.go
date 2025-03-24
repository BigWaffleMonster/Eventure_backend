package views

// @description user information
type UserInfo struct {
	Email    string  `json:"email"`
	Password string `json:"password"`
}

// @description login data
type LoginInfo struct {
	Email    string  `json:"email"`
	Password *string `json:"password"`
}