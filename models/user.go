package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	UserName string
	Password string
	UserType string
	RolesID  uint
}
