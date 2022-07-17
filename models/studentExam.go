package models

import "gorm.io/gorm"

type StudentExam struct {
	gorm.Model
	ID             uint `gorm:"primary_key:autoIncrement"`
	StudentID      uint `gorm:"foreignKey"`
	ExamID         uint `gorm:"foreignKey"`
	SessionYearID  uint `gorm:"foreignKey"`
	ClassID        uint
	SubjectID      uint
	ExamType       string
	InternalMarks  uint
	TheoryMarks    uint
	PracticalMarks uint
	MaximumMarks   uint
	Exam           Exam
	Subject        Subject
}
