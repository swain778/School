package models

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	Name        string `gorm:"not null;default:null"`
	DOB         string `gorm:"not null;default:null"`
	JoiningDate string `gorm:"not null;default:null"`
	Session     string `gorm:"not null;default:null"`
	Aadharno    string `gorm:"not null;default:null"`
	StaffType   string `gorm:"not null;default:null"`
	Salary      Salary
	BankDetail  BankDetail
}
