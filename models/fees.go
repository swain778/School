package models

import (
	"time"

	"gorm.io/gorm"
)

type Fees struct {
	gorm.Model
	ID        uint `gorm:"primary_key:autoIncrement"`
	StudentID uint `gorm:"foreignKey"`
	Session   string
	FeesPaid  string
	TotalFees string
	Pending   string
	ClassID   uint      `gorm:"foreignKey"`
	CreatedAt time.Time `gorm:"autoUpdate"`
	UpdatedAt time.Time `gorm:"autoUpdate"`
}
