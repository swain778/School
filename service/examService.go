package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ExamService struct {
	db *gorm.DB
}

func NewExamService() *ExamService {
	examService := new(ExamService)
	examService.db = config.GetDB()
	return examService
}

func (e *ExamService) CreateExam(exam *models.Exam) (*models.Exam, error) {
	err := e.db.Where(models.Exam{
		SubjectID: exam.SubjectID,
		ClassID:   exam.ClassID,
	}).FirstOrCreate(&exam).Error
	if err != nil {
		return nil, errors.New("can't create exam")
	}
	return exam, nil
}

func (e *ExamService) DeleteExam(examID string) (bool, error) {
	exam := &models.Exam{}
	err := e.db.Where("id =?", examID).Delete(exam).Error
	if err != nil {
		return false, errors.New("can't delete exam")
	}
	return true, nil
}

func (e *ExamService) GetExam(examID string) (*models.Exam, error) {
	exam := &models.Exam{}
	err := e.db.Model(&models.Exam{}).Preload(clause.Associations).First(exam, "id=?", examID).Error
	if err != nil {
		return nil, errors.New("can't get exam")
	}
	return exam, nil
}

func (e *ExamService) GetsExam() (*[]models.Exam, error) {
	exam := &[]models.Exam{}
	err := e.db.Model(&models.Exam{}).Find(exam).Error
	if err != nil {
		return nil, errors.New("can't get exam id")
	}
	return exam, nil
}

func (e *ExamService) UpdateExam(exam *models.Exam) (bool, error) {
	err := e.db.Where("id=?", exam.ID).Updates(&exam).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}
