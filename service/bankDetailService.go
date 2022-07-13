package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
)

type BankDetail struct {
	db *gorm.DB
}

func NewBankDetailervice() *BankDetail {
	BankDetail := new(BankDetail)
	BankDetail.db = config.GetDB()
	return BankDetail
}

func (b *BankDetail) CreateBankDetail(bankDetail *models.BankDetail) (*models.BankDetail, error) {
	err := b.db.Where(models.BankDetail{
		Bank:        bankDetail.Bank,
		BankAccount: bankDetail.BankAccount,
		Name:        bankDetail.Name,
	}).FirstOrCreate(&bankDetail).Error
	if err != nil {
		return nil, errors.New("can't enter bank details")
	}
	return bankDetail, nil
}

func (b *BankDetail) DeleteBankDetail(BankDetailID string) (bool, error) {
	err := b.db.Where("id=?", BankDetailID).Delete(&models.BankDetail{}).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (b *BankDetail) GetsBankDetail(staffID uint) (*[]models.BankDetail, error) {
	BankDetail := &[]models.BankDetail{}
	err := b.db.Where("staff_id=?", staffID).Find(BankDetail).Error
	if err != nil {
		return nil, errors.New("can't get bank details")
	}
	return BankDetail, nil
}
