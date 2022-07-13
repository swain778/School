package models

import "gorm.io/gorm"

type SessionYear struct {
	gorm.Model
	ID          uint `gorm:"primary_key:autoIncrement"`
	SessionYear string
	StudentExam StudentExam
}
