package models

import "gorm.io/gorm"

type StudentSession struct {
	gorm.Model
	ID           uint `gorm:"primary_key:autoIncrement"`
	StudentID    uint `gorm:"foreignKey"`
	ClassID      uint `gorm:"foreignKey"`
	Section      string
	SessionYear  string
	Semister     string
	Incharge     string
	Student      Student
	StudentClass []StudentClass
	StudentExam  []StudentExam
}
