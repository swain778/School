package models

import (
	"time"

	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	ID      uint   `gorm:"primary_key:autoIncrement"`
	Subject string `gorm:"not null;default:null"`
	//Exam          Exam
	CreatedAt time.Time `gorm:"autoUpdate"`
	UpdatedAt time.Time `gorm:"autoUpdate"`
	//Class     Class
	Homework Homework
}
