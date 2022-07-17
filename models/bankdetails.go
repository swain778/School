package models

import "gorm.io/gorm"

type BankDetail struct {
	gorm.Model
	StaffID     uint   `gorm:"foreignKey"`
	Name        string `gorm:"not null;default:null"`
	Bank        string `gorm:"not null;default:null"`
	BankAccount string `gorm:"not null;default:null"`
	IFSC        string `gorm:"not null;default:null"`
	BranchCode  string `gorm:"not null;default:null"`
	IsDefault   bool   `gorm:"default:true"`
}
