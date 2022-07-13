package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID     uint
	Role   string
	Status string
	User   []*Login
}
