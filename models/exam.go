package models

import (
	"time"

	"gorm.io/gorm"
)

type Exam struct {
	gorm.Model
	ID          uint `gorm:"primary_key:autoIncrement"`
	SubjectID   uint `gorm:"foreignKey"`
	ClassID     uint `gorm:"foreignKey"`
	StudentExam []StudentExam
	//Subject     Subject
	CreatedAt time.Time `gorm:"autoUpdate"`
	UpdatedAt time.Time `gorm:"autoUpdate"`
}
