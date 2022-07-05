package models

import (
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ID                uint   `gorm:"primary_key:autoIncrement"`
	FirstName         string `gorm:"not null"`
	LastName          string `gorm:"not null"`
	Department        string
	DOB               string
	JoiningAt         string
	Status            string
	TeacherAttendence []TeacherAttendence
	ClassTeacher      []ClassTeacher
	CreatedAt         time.Time `gorm:"autoUpdate"`
	UpdatedAt         time.Time `gorm:"autoUpdate"`
}
