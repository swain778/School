package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClassTeacherService struct {
	db *gorm.DB
}

func NewClassTeacherService() *ClassTeacherService {
	classTeacherService := new(ClassTeacherService)
	classTeacherService.db = config.GetDB()
	return classTeacherService
}

func (t *ClassTeacherService) CreateClassTeacher(classTeacher *models.ClassTeacher) (*models.ClassTeacher, error) {
	err := t.db.Where(models.ClassTeacher{
		ClassID:   classTeacher.ClassID,
		TeacherID: classTeacher.TeacherID,
	}).FirstOrCreate(&classTeacher).Error
	if err != nil {
		return nil, errors.New("can't Create Class Teacher")
	}
	return classTeacher, nil
}

func (t *ClassTeacherService) DeleteClassTeacher(classTeacherID string) (bool, error) {
	classTeacher := &models.ClassTeacher{}
	err := t.db.Where("id = ?", classTeacherID).Delete(classTeacher).Error
	if err != nil {
		return false, errors.New("can't delete Class Teacher")
	}
	return true, nil
}

func (t *ClassTeacherService) UpdateClassTeacher(classTeacher *models.ClassTeacher) (bool, error) {
	err := t.db.Where("id =?", classTeacher.ID).Updates(&classTeacher).Error
	if err != nil {
		return false, errors.New("can't update class teacher")
	}
	return true, nil
}

func (t *ClassTeacherService) GetClassTeacher(classTeacherID string) (bool, error) {
	classTeacher := &models.ClassTeacher{}
	err := t.db.Model(&models.ClassTeacher{}).Preload(clause.Associations).First(classTeacher, "id=?", classTeacherID).Error
	if err != nil {
		return false, errors.New("can't get class ID")
	}
	return true, nil
}

func (t *ClassTeacherService) GetsClassTeacher() (*[]models.ClassTeacher, error) {
	classTeacher := &[]models.ClassTeacher{}
	err := t.db.Model(&[]models.ClassTeacher{}).Find(classTeacher).Error
	if err != nil {
		return nil, errors.New("can't get class teacher")
	}
	return classTeacher, nil
}
