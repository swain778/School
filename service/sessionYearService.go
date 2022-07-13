package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SessionYear struct {
	db *gorm.DB
}

func NewSessionYear() *SessionYear {
	SessionYear := new(SessionYear)
	SessionYear.db = config.GetDB()
	return SessionYear
}

func (s *SessionYear) CreateSessionYear(session *models.SessionYear) (*models.SessionYear, error) {
	err := s.db.Where(models.SessionYear{
		SessionYear: session.SessionYear,
	}).FirstOrCreate(&session).Error
	if err != nil {
		return nil, errors.New("can't create student seassion")
	}
	return session, nil
}

func (s *SessionYear) DeleteSessionYear(sessionID string) (bool, error) {
	session := &models.SessionYear{}
	err := s.db.Where("id=?", sessionID).Delete(session).Error
	if err != nil {
		return false, errors.New("can't delete session")
	}
	return true, nil
}

func (s *SessionYear) UpdateSessionYear(session *models.SessionYear) (bool, error) {
	err := s.db.Where("id=", session.ID).First(&session).Error
	if err != nil {
		return false, errors.New("can't update session")
	}
	return true, nil
}

func (s *SessionYear) GetSessionYear(sessionID string) (*models.SessionYear, error) {
	session := &models.SessionYear{}
	err := s.db.Model(&models.SessionYear{}).Preload(clause.Associations).First(session, "id=?", sessionID).Error
	if err != nil {
		return nil, errors.New("can't get seassion Id")
	}
	return session, nil
}

func (s *SessionYear) GetsSessionYear() (*[]models.SessionYear, error) {
	session := &[]models.SessionYear{}
	err := s.db.Model(&[]models.SessionYear{}).Find(session).Error
	if err != nil {
		return nil, errors.New("can't gert student session")
	}
	return session, nil
}
