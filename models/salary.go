package models

import "gorm.io/gorm"

type Salary struct {
	gorm.Model
	Basic        int    `gorm:"not null;default:null"`
	Bonus        int    `gorm:"not null;default:null"`
	PF           int    `gorm:"not null;default:null"`
	Tax          int    `gorm:"not null;default:null"`
	NetAmount    int    `gorm:"not null;default:null"`
	Month        string `gorm:"not null;default:null"`
	StaffID      uint   `gorm:"foreignKey"`
	BankDetailID uint   `gorm:"foreignKey"`
	BankDetail   *BankDetail
}
