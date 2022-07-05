package models

import "gorm.io/gorm"

type ClassTeacher struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	ClassID   uint `gorm:"foreignKey"`
	TeacherID uint `gorm:"foreignKey"`
}
