package models

import (
	"time"

	"gorm.io/gorm"
)

type TeacherAttendence struct {
	gorm.Model
	ID                uint `gorm:"primary_key:autoIncrement"`
	TeacherAttendence string
	TeacherID         uint      `gorm:"foreignKey"`
	CreatedAt         time.Time `gorm:"autoUpdate"`
	UpdatedAt         time.Time `gorm:"autoUpdate"`
}
