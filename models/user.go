package models

import "gorm.io/gorm"

type RoleType string

const (
	AdminRole RoleType = "admin"
	UserRole RoleType = "user"
	ViewerRole RoleType = "viewer"
)

type Role struct {
	gorm.Model
	Name RoleType `gorm:"unique;not null" json:"name"`
}


type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null" form:"email" json:"email"`
	Phone    string `form:"phone" json:"phone"`
	Password string `form:"password" json:"password"`
	RoleID  uint  `json:"role_id"`
	Role  Role `gorm:"foreignkey:RoleID" json:"role"`
}
