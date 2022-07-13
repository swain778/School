package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StudentClassService struct {
	db *gorm.DB
}

func NewStudentClassService() *StudentClassService {
	studentClass := new(StudentClassService)
	studentClass.db = config.GetDB()
	return studentClass
}

func (s *StudentClassService) CreateStudentClassService(studentClass *models.StudentClass) (*models.StudentClass, error) {
	err := s.db.Where(models.StudentClass{
		StudentID:     studentClass.StudentID,
		ClassID:       studentClass.ClassID,
		SessionYearID: studentClass.SessionYearID,
	}).FirstOrCreate(&studentClass).Error
	if err != nil {
		return nil, errors.New("can't create student service")
	}
	return studentClass, nil
}

func (s *StudentClassService) DeleteStudentClassService(studentClassID string) (bool, error) {
	studentClass := &models.StudentClass{}
	err := s.db.Where("id =?", studentClassID).Delete(studentClass).Error
	if err != nil {
		return false, errors.New("can't delete student class")
	}
	return true, nil
}

func (s *StudentClassService) UpdateStudentClassService(studentClass *models.StudentClass) (*models.StudentClass, error) {
	err := s.db.Where("id=?", studentClass.ID).Updates(&studentClass).Error
	if err != nil {
		return nil, errors.New("can't update student class")
	}
	return studentClass, nil
}

func (s *StudentClassService) GetStudentClassService(studentClassID string) (bool, error) {
	studentClass := &models.StudentClass{}
	err := s.db.Model(&models.StudentClass{}).Preload(clause.Associations).First(studentClass, "id=?", studentClassID)
	if err != nil {
		return false, errors.New("can't get student details")
	}
	return true, nil
}

func (s *StudentClassService) GetsStudentClassService() (*[]models.StudentClass, error) {
	studentClass := &[]models.StudentClass{}
	err := s.db.Model(&[]models.StudentClass{}).Find(studentClass).Error
	if err != nil {
		return nil, errors.New("can't gets student")
	}
	return studentClass, nil
}
