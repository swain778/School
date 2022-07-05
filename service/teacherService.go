package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Teacher struct {
	db *gorm.DB
}

func NewTeacher() *Teacher {
	teacher := new(Teacher)
	teacher.db = config.GetDB()
	return teacher
}

func (t *Teacher) CreateTeacher(teacher *models.Teacher) (*models.Teacher, error) {
	err := t.db.Where(models.Teacher{
		FirstName:  teacher.FirstName,
		LastName:   teacher.LastName,
		Department: teacher.Department,
		DOB:        teacher.DOB,
		JoiningAt:  teacher.JoiningAt,
		Status:     teacher.Status,
	}).FirstOrCreate(&teacher).Error
	if err != nil {
		return nil, errors.New("can't create teacher")
	}
	return teacher, nil
}

func (t *Teacher) DeleteTeacher(teacherID string) (bool, error) {
	teacher := &models.Teacher{}
	err := t.db.Where("id=?", teacherID).Delete(teacher).Error
	if err != nil {
		return false, errors.New("can't delete teacher")
	}
	return true, nil
}

func (t *Teacher) UpdateTeacher(teacher *models.Teacher) (*models.Teacher, error) {
	err := t.db.Where("id=?", teacher.ID).Updates(&teacher).Error
	if err != nil {
		return nil, errors.New("can't update teacher")
	}
	return teacher, nil
}

func (t *Teacher) GetTeacher(teacherID string) (*models.Teacher, error) {
	teacher := &models.Teacher{}
	err := t.db.Model(&models.Teacher{}).Preload(clause.Associations).First(teacher, "id=?", teacherID).Error
	if err != nil {
		return nil, errors.New("can't get teacher")
	}
	return teacher, nil
}

func (t *Teacher) GetsTeacher() (*[]models.Teacher, error) {
	teacher := &[]models.Teacher{}
	err := t.db.Model(&[]models.Teacher{}).Find(teacher).Error
	if err != nil {
		return nil, errors.New("can't get teachers details")
	}
	return teacher, nil
}
