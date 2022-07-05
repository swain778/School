package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID                uint `gorm:"primary_key:autoIncrement:"`
	Name              string
	Address           string
	DOB               string
	Father_Name       string
	Mother_Name       string
	Status            string
	StudentAttendence []StudentAttendence
	StudentClass      []StudentClass
	Complain          []Complain
	StudentExam       []StudentExam
	CreatedAt         time.Time `gorm:"autoUpdate"`
	UpdatedAt         time.Time `gorm:"autoUpdate"`
}
