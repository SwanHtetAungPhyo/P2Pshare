package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func NewUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}
