package models

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ID           uint   `gorm:"primary_key:autoIncrement"`
	Class        string `gorm:"unique:not null;default:null"`
	ClassTeacher []ClassTeacher
	Subject      []Subject
	StudentClass []StudentClass
	Fees         []Fees
	CreatedAt    time.Time `gorm:"autoUpdate"`
	UpdatedAt    time.Time `gorm:"autoUpdate"`
}
