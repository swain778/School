package models

import (
	"time"

	"gorm.io/gorm"
)

type Fees struct {
	gorm.Model
	ID        uint   `gorm:"primary_key:autoIncrement"`
	StudentID uint   `gorm:"foreignKey"`
	Session   string `gorm:"not null;default:0"`
	FeesPaid  uint   `gorm:"not null;default:0"`
	TotalFees uint   `gorm:"not null;default:0"`
	Pending   uint   `gorm:"not null;default:0"`
	Month     string
	ClassID   uint      `gorm:"foreignKey"`
	CreatedAt time.Time `gorm:"autoUpdate"`
	UpdatedAt time.Time `gorm:"autoUpdate"`
}
