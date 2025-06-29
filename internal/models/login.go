package models

import "github.com/go-playground/validator"

type UserLogin struct {
	Username string `json:"username"  validate:"required"`
	Password string `json:"password"  validate:"required,min=8,max=100"`
}

func (l UserLogin) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserLoginResponse struct {
	UserID       uint   `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
