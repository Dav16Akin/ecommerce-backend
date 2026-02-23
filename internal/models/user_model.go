package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Email       string
	Password    string `json:"-"`
	PhoneNumber string
	Orders      []Order `gorm:"foreignKey:UserID"`
}

type CreateUserRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type UpdateUserRequest struct {
	Name        *string `json:"name" binding:"omitempty"`
	Email       *string `json:"email" binding:"omitempty"`
	PhoneNumber *string `json:"phone_number" binding:"omitempty"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type Order struct {
	gorm.Model
	UserID uint
	User   User
}
