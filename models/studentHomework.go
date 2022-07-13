package models

import "time"

type StudentHomework struct {
	ID         uint
	HomeworkID uint
	StudentID  uint
	CreatedAt  time.Time
	DeletedAt  time.Time
	UpdatedAt  time.Time
}
