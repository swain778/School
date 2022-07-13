package models

import "gorm.io/gorm"

type Homework struct {
	gorm.Model
	SubjectID      uint
	StudentID      uint
	TeacherID      uint
	Class          string
	Homework       string
	SubmissionDate string
}
