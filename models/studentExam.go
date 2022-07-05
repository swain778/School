package models

import "gorm.io/gorm"

type StudentExam struct {
	gorm.Model
	ID               uint `gorm:"primary_key:autoIncrement"`
	StudentID        uint `gorm:"foreignKey"`
	ExamID           uint `gorm:"foreignKey"`
	StudentSessionID uint `gorm:"foreignKey"`
	Exam             Exam
}
