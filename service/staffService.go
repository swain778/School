package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
)

type StaffService struct {
	db gorm.DB
}

func NewStaffService() *StaffService {
	staffService := new(StaffService)
	staffService.db = *config.GetDB()
	return staffService
}

func (f *StaffService) CreateStaff(staffService *models.Staff) (*models.Staff, error) {
	err := f.db.Where(&models.Staff{
		StaffType: staffService.StaffType,
		Name:      staffService.Name,
		Aadharno:  staffService.Aadharno,
	}).FirstOrCreate(&staffService).Error
	if err != nil {
		return nil, errors.New("can't create staff")
	}
	return staffService, nil
}

func (f *StaffService) DeleteStaff(staffServiceID string) (bool, error) {
	err := f.db.Where("id=?", staffServiceID).Delete(&models.Staff{}).Error
	if err != nil {
		return false, errors.New("can't delete staff")
	}
	return true, nil
}

func (f *StaffService) GetStaff(staffType string) ([]*models.Staff, error) {
	staffData := []*models.Staff{}
	err := f.db.Where("staff_type ILIKE ?", staffType).Find(&staffData).Error
	if err != nil {
		return nil, errors.New("can't get staff")
	}
	return staffData, nil
}
func (f *StaffService) GetStaffByID(staffType string, staffID string) (*models.Staff, error) {
	staffData := &models.Staff{}
	err := f.db.Where("staff_type ILIKE ?", staffType).Where("id = ?", staffID).Find(&staffData).Error
	if err != nil {
		return nil, errors.New("can't get staff")
	}
	return staffData, nil
}
