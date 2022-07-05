package models

import (
	"time"

	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	ID        uint `gorm:"primary_key:autoIncrement"`
	Subject   string
	ClassID   uint      `gorm:"foreignKey"`
	CreatedAt time.Time `gorm:"autoUpdate"`
	UpdatedAt time.Time `gorm:"autoUpdate"`
}
