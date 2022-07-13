package models

import (
	"time"

	"gorm.io/gorm"
)

type Complain struct {
	gorm.Model
	ID        uint   `gorm:"primary_key:autoIncrement"`
	Complain  string `gorm:"not null"`
	StudentID uint   `gorm:"foreignKey"`
	Student   Student
	CreatedAt time.Time `gorm:"autoUpdate"`
	UpdatedAt time.Time `gorm:"autoUpdate"`
}
