package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FeesService struct {
	db *gorm.DB
}

func NewFeeService() *FeesService {
	feesService := new(FeesService)
	feesService.db = config.GetDB()
	return feesService
}

func (f *FeesService) CreateFees(fees *models.Fees) (*models.Fees, error) {

	err := f.db.
		Where(models.Fees{
			ClassID:   fees.ClassID,
			StudentID: fees.StudentID,
		}).
		FirstOrCreate(fees).Error

	if err != nil {
		return nil, errors.New("can't create fees")
	}
	return fees, nil
}

func (f *FeesService) DeleteFees(feesID string) (bool, error) {
	fees := &models.Fees{}
	err := f.db.Where("id=?", feesID).Delete(fees).Error
	if err != nil {
		return false, errors.New("can't delete fees")
	}
	return true, nil
}

func (f *FeesService) UpdateFees(fees *models.Fees) (bool, error) {
	err := f.db.Where("id=", fees.ID).Updates(&fees)
	if err != nil {
		return false, errors.New("can't update fees")
	}
	return true, nil
}

func (f *FeesService) GetFees(feesID string) (*models.Fees, error) {
	fees := &models.Fees{}
	err := f.db.Model(&models.Fees{}).Preload(clause.Associations).First(fees, "id=?", feesID)
	if err != nil {
		return nil, errors.New("can't get ID")
	}
	return fees, nil
}

func (f *FeesService) GetsFees() (*[]models.Fees, error) {
	fees := &[]models.Fees{}
	err := f.db.Model(&[]models.Fees{}).Find(fees).Error
	if err != nil {
		return nil, errors.New("can't gets fees")
	}
	return fees, nil
}
