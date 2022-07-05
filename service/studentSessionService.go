package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StudentSession struct {
	db *gorm.DB
}

func NewStudentSession() *StudentSession {
	studentSession := new(StudentSession)
	studentSession.db = config.GetDB()
	return studentSession
}

func (s *StudentSession) CreateStudentSession(session *models.StudentSession) (*models.StudentSession, error) {
	err := s.db.Where(models.StudentSession{
		ClassID:     session.ClassID,
		StudentID:   session.StudentID,
		Section:     session.Section,
		SessionYear: session.SessionYear,
		Semister:    session.Semister,
		Incharge:    session.Incharge,
	}).FirstOrCreate(&session).Error
	if err != nil {
		return nil, errors.New("can't create student seassion")
	}
	return session, nil
}

func (s *StudentSession) DeleteStudentSession(sessionID string) (bool, error) {
	session := &models.StudentSession{}
	err := s.db.Where("id=?", sessionID).Delete(session).Error
	if err != nil {
		return false, errors.New("can't delete session")
	}
	return true, nil
}

func (s *StudentSession) UpdateStudentSession(session *models.StudentSession) (bool, error) {
	err := s.db.Where("id=", session.ID).First(&session).Error
	if err != nil {
		return false, errors.New("can't update session")
	}
	return true, nil
}

func (s *StudentSession) GetStudentSession(sessionID string) (*models.StudentSession, error) {
	session := &models.StudentSession{}
	err := s.db.Model(&models.StudentSession{}).Preload(clause.Associations).First(session, "id=?", sessionID).Error
	if err != nil {
		return nil, errors.New("can't get seassion Id")
	}
	return session, nil
}

func (s *StudentSession) GetsStudentSession() (*[]models.StudentSession, error) {
	session := &[]models.StudentSession{}
	err := s.db.Model(&[]models.StudentSession{}).Find(session).Error
	if err != nil {
		return nil, errors.New("can't gert student session")
	}
	return session, nil
}
