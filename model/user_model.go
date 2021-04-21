package model

type RegisterUserRequest struct {
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfileRequest struct {
	ID      string
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Profile string `json:"profile"`
}
