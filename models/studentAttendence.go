package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentAttendence struct {
	gorm.Model
	Id         uint `gorm:"primary_key:autoIncrement"`
	Attendence string
	StudentID  uint `gorm:"foreignKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
