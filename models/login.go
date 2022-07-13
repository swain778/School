package models

import (
	"gorm.io/gorm"
)

type UserType string

const (
	AdminUser   UserType = "Admin"
	TeacherUser UserType = "Teacher"
	StudentUser UserType = "Student"
)

type Login struct {
	gorm.Model
	UserName   string `gorm:"unique"`
	Password   string
	UserType   UserType
	UserTypeID uint
	RoleID     uint
	Token      string
	Role       *Role //`gorm:"foreignKey:RoleID"`
}
