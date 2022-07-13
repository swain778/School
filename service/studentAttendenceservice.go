package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StudentAttendence struct {
	db *gorm.DB
}

func NewAttendenceService() *StudentAttendence {
	studentAttendence := new(StudentAttendence)
	studentAttendence.db = config.GetDB()
	return studentAttendence
}

func (a *StudentAttendence) CreateAttendence(attendence *models.StudentAttendence) (*models.StudentAttendence, error) {
	err := a.db.Where(&models.StudentAttendence{
		StudentID:  attendence.StudentID,
		Attendence: attendence.Attendence,
		Date:       attendence.Date,
	}).FirstOrCreate(&attendence).Error
	if err != nil {
		return nil, errors.New("can't create attendence")
	}
	return attendence, nil
}

func (a *StudentAttendence) DeleteAttendence(attendenceID string) (bool, error) {
	attendence := &models.StudentAttendence{}
	err := a.db.Where("id=?", attendenceID).Delete(attendence).Error
	if err != nil {
		return false, errors.New("can't delete attendence")
	}
	return true, nil
}

func (a *StudentAttendence) UpdateAttendence(attendence *models.StudentAttendence) (bool, error) {
	err := a.db.Where("id=?", attendence.Id).Updates(&attendence)
	if err != nil {
		return false, errors.New("can't update attendence")
	}
	return true, nil
}

func (a *StudentAttendence) GetAttendence(attendenceID string) (*models.StudentAttendence, error) {
	attendence := &models.StudentAttendence{}
	err := a.db.Model(&models.StudentAttendence{}).Preload(clause.Associations).First(attendence, "id=?", attendenceID)
	if err != nil {
		return nil, errors.New("can't get id")
	}
	return attendence, nil
}

func (a *StudentAttendence) GetsAttendence() (*[]models.StudentAttendence, error) {
	attendence := &[]models.StudentAttendence{}
	err := a.db.Model(&[]models.StudentAttendence{}).Find(attendence).Error
	if err != nil {
		return nil, errors.New("can't get attendence")
	}
	return attendence, nil
}
