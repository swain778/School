package models

import "gorm.io/gorm"

type ClassSubject struct {
	gorm.Model
	ClassID   uint
	SubjectID uint
}
