package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TeacherAttendence struct {
	db *gorm.DB
}

func NewTeacherAttendence() *TeacherAttendence {
	teacherAttendence := new(TeacherAttendence)
	teacherAttendence.db = config.GetDB()
	return teacherAttendence
}

func (t *TeacherAttendence) CreateTeacherAttendence(teacherAttendence *models.TeacherAttendence) (*models.TeacherAttendence, error) {
	err := t.db.Where(&models.TeacherAttendence{
		TeacherID:         teacherAttendence.TeacherID,
		TeacherAttendence: teacherAttendence.TeacherAttendence,
		Date:              teacherAttendence.Date,
	}).FirstOrCreate(&teacherAttendence).Error
	if err != nil {
		return nil, errors.New("can't create teacher attendence")
	}
	return teacherAttendence, nil
}

func (t *TeacherAttendence) DeleteTeacherAttendence(teacherAttendenceID string) (bool, error) {
	teacherAttendence := &models.TeacherAttendence{}
	err := t.db.Where("id=?", teacherAttendenceID).Delete(teacherAttendence).Error
	if err != nil {
		return false, errors.New("can't delete teacherAttendence")
	}
	return true, nil
}

func (t *TeacherAttendence) UpdateTeacherAttendence(teacherAttendence *models.TeacherAttendence) (*models.TeacherAttendence, error) {
	err := t.db.Where("id=?", teacherAttendence.ID).Updates(&teacherAttendence).Error
	if err != nil {
		return nil, errors.New("can't update teacher attendence")
	}
	return teacherAttendence, nil
}

func (t *TeacherAttendence) GetTeacherAttendence(teacherAttendenceID string) (*models.TeacherAttendence, error) {
	teacherAttendence := &models.TeacherAttendence{}
	err := t.db.Model(&models.TeacherAttendence{}).Preload(clause.Associations).First(teacherAttendence, "id=?", teacherAttendenceID).Error
	if err != nil {
		return nil, errors.New("can't get teacher attendence")
	}
	return teacherAttendence, nil
}

func (t *TeacherAttendence) GetsTeacherAttendence() (*[]models.TeacherAttendence, error) {
	teacherAttendence := &[]models.TeacherAttendence{}
	err := t.db.Model(&[]models.TeacherAttendence{}).Find(teacherAttendence).Error
	if err != nil {
		return nil, errors.New("can't gets teacher Attendence")
	}
	return teacherAttendence, nil
}
