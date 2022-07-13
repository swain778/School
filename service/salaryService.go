package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
)

type SalaryService struct {
	db *gorm.DB
}

func NewSalaryService() *SalaryService {
	salaryService := new(SalaryService)
	salaryService.db = config.GetDB()
	return salaryService
}

func (s *SalaryService) CreateSalaryService(salaryService *models.Salary) (*models.Salary, error) {
	err := s.db.Where(models.Salary{
		StaffID: salaryService.StaffID,
		Month:   salaryService.Month,
	}).FirstOrCreate(salaryService).Error
	if err != nil {
		return nil, errors.New("can't create salary")
	}
	return salaryService, nil
}

func (s *SalaryService) DeleteSalaryService(salaryID string) (bool, error) {
	salary := &models.Salary{}
	err := s.db.Debug().Where("id=?", salaryID).Delete(salary).Error
	if err != nil {
		return false, errors.New("can't delete salary ")
	}
	return true, nil
}
