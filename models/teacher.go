package models

import (
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ID                uint   `gorm:"primary_key:autoIncrement"`
	FirstName         string `gorm:"not null;default:null"`
	LastName          string `gorm:"not null;default:null"`
	Department        string `gorm:"not null;default:null"`
	DOB               string `gorm:"not null;default:null"`
	JoiningAt         string `gorm:"not null;default:null"`
	Status            string `gorm:"not null;default:null"`
	TeacherAttendence []TeacherAttendence
	ClassTeacher      []ClassTeacher
	CreatedAt         time.Time `gorm:"autoUpdate"`
	UpdatedAt         time.Time `gorm:"autoUpdate"`
}
