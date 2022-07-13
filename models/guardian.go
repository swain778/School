package models

import (
	"time"

	"gorm.io/gorm"
)

type GuardianType string

const (
	GuardianFather   GuardianType = "Father"
	GuardianMother   GuardianType = "Mother"
	GuardianSibling  GuardianType = "Sibling"
	GuardianRelative GuardianType = "Relative"
)

type Guardian struct {
	ID           uint
	GuardianType string
	GuardianName string
	StudentID    uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	gorm.DeletedAt
}
