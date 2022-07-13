package models

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	ResultType string
	RollNumber string
	Subject    Subject
	Student    Student
	Class      Class
	Teacher    Teacher
}
