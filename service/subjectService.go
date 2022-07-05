package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubjectService struct {
	db *gorm.DB
}

func NewSubjectService() *SubjectService {
	subject := new(SubjectService)
	subject.db = config.GetDB()
	return subject
}

func (s *SubjectService) CreateSubject(subject *models.Subject) (*models.Subject, error) {
	err := s.db.Where(&models.Subject{
		ClassID: subject.ClassID,
		Subject: subject.Subject,
	}).FirstOrCreate(&subject).Error

	if err != nil {
		return nil, errors.New("can't create subject")
	}
	return subject, nil
}

func (s *SubjectService) DeleteSubject(subjectID string) (bool, error) {
	subject := &models.Subject{}
	err := s.db.Where("id=?", subjectID).Delete(subject).Error
	if err != nil {
		return false, errors.New("can't delete subject")
	}
	return true, nil
}

func (s *SubjectService) UpdateSubject(subject *models.Subject) (bool, error) {
	err := s.db.Where("id=?", subject.ID).Updates(&subject).Error
	if err != nil {
		return false, errors.New("can't update subject")
	}
	return true, nil
}

func (s *SubjectService) GetSubject(subjectID string) (*models.Subject, error) {
	subject := &models.Subject{}
	err := s.db.Model(&models.Subject{}).Preload(clause.Associations).First(subject, "id=?", subjectID).Error
	if err != nil {
		return nil, errors.New("can't get subject id")
	}
	return subject, nil
}

func (s *SubjectService) GetsSubject() (*[]models.Subject, error) {
	subject := &[]models.Subject{}
	err := s.db.Model(&[]models.Subject{}).Find(subject).Error
	if err != nil {
		return nil, errors.New("can't get subjects")
	}
	return subject, nil
}
