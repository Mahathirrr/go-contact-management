package models

type User struct {
	Username string  `json:"username" db:"username"`
	Password string  `json:"password,omitempty" db:"password"`
	Name     string  `json:"name" db:"name"`
	Token    *string `json:"token,omitempty" db:"token"`
}

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Name     string `json:"name" validate:"required,max=100"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type UserUpdateRequest struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,max=100"`
	Password *string `json:"password,omitempty" validate:"omitempty,max=100"`
}

type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
