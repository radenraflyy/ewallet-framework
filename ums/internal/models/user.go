package models

import (
	"time"

	"github.com/go-playground/validator"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"unique;not null" validate:"required"`
	Email       string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	PhoneNumber string    `json:"phone_number" gorm:"unique;not null" validate:"required"`
	Address     string    `json:"address" gorm:"not null"`
	DateOfBirth string    `json:"date_of_birth" gorm:"not null"`
	FullName    string    `json:"full_name" gorm:"not null" validate:"required"`
	Password    string    `json:"password,omitempty" gorm:"not null" validate:"required,min=8,max=100"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (*User) TableName() string {
	return "users"
}
func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	UserID             uint      `json:"user_id" gorm:"not null" validate:"required"`
	Token              string    `json:"token" gorm:"not null" validate:"required"`
	RefreshToken       string    `json:"refresh_token" gorm:"not null" validate:"required"`
	TokenExpiry        time.Time `json:"token_expiry" gorm:"not null" validate:"required"`
	RefreshTokenExpiry time.Time `json:"refresh_token_expiry" gorm:"not null" validate:"required"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (*UserSession) TableName() string {
	return "users_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
