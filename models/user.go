package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null" form:"email" json:"email"`
	Phone    string `form:"phone" json:"phone"`
	Password string `form:"password" json:"password"`
}
