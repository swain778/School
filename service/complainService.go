package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ComplainService struct {
	db *gorm.DB
}

func NewComplainService() *ComplainService {
	complainService := new(ComplainService)
	complainService.db = config.GetDB()
	return complainService
}

func (c *ComplainService) CreateComplain(complain *models.Complain) (*models.Complain, error) {

	err := c.db.Where(&models.Complain{
		StudentID: complain.StudentID,
		Complain:  complain.Complain,
		Student: models.Student{
			ID: complain.StudentID,
		},
	}).FirstOrCreate(&complain).Error

	if err != nil {
		return nil, errors.New("can't create complain")
	}
	return complain, nil
}

func (c *ComplainService) DeleteComplain(complainID string) (bool, error) {
	complain := &models.Complain{}
	err := c.db.Where("id =?", complainID).Delete(complain).Error
	if err != nil {
		return false, errors.New("can't delete complain")
	}
	return true, nil
}

func (c *ComplainService) GetComplain(complainID string) (*models.Complain, error) {
	complain := &models.Complain{}
	err := c.db.Model(&models.Complain{}).Preload(clause.Associations).First(complain, "id=?", complainID).Error
	if err != nil {
		return nil, errors.New("can't get complain")
	}
	return complain, nil
}

func (c *ComplainService) GetsComplain() (*[]models.Complain, error) {
	complain := &[]models.Complain{}
	err := c.db.Model(&[]models.Complain{}).Find(complain).Error
	if err != nil {
		return nil, errors.New("can't get complain")
	}
	return complain, nil
}

func (c *ComplainService) UpdateComplain(complain *models.Complain) (bool, error) {
	err := c.db.Where("id=?", complain.ID).Updates(&complain).Error
	if err != nil {
		return false, errors.New("can't update complain")
	}
	return true, nil
}
