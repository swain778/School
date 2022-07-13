package models

import "gorm.io/gorm"

type StudentSession struct {
	gorm.Model
	ID           uint   `gorm:"primary_key:autoIncrement"`
	StudentID    uint   `gorm:"foreignKey"`
	ClassID      uint   `gorm:"foreignKey"`
	Section      string `gorm:"not null;default:null"`
	SessionYear  string `gorm:"not null;default:null"`
	Semister     string `gorm:"not null;default:null"`
	Incharge     string `gorm:"not null;default:null"`
	Student      Student
	StudentClass []StudentClass
	StudentExam  []StudentExam
}
