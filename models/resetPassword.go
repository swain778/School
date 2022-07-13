package models

import "time"

type ResetPassword struct {
	ID               uint
	UserID           uint
	ResetKey         string
	ValidTill        time.Time `gorm:"autoIncrement;default:null"`
	IsUsed           bool
	UpdatePasswordID uint
}
