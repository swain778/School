package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
)

type GuardianService struct {
	db *gorm.DB
}

func NewGuardianService() *GuardianService {
	guardianService := new(GuardianService)
	guardianService.db = config.GetDB()
	return guardianService
}

func (g *GuardianService) CreateGuardian(guardian *models.Guardian) (*models.Guardian, error) {
	err := g.db.Where(models.Guardian{
		StudentID:    guardian.StudentID,
		GuardianName: guardian.GuardianName,
	}).FirstOrCreate(&guardian).Error
	if err != nil {
		return nil, errors.New("can't create guardian")
	}
	return guardian, nil
}

func (g *GuardianService) DeleteGuardian(guardianID string) (bool, error) {

	err := g.db.Where("id=?", guardianID).Delete(&models.Guardian{}).Error
	if err != nil {
		return false, errors.New("can't delete guardian")
	}
	return true, nil
}
