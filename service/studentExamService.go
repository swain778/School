package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StudentExamService struct {
	db *gorm.DB
}

func NewStudentExamService() *StudentExamService {
	examService := new(StudentExamService)
	examService.db = config.GetDB()
	return examService
}

func (e *StudentExamService) SaveStudentExamInfo(exam *models.StudentExam) (*models.StudentExam, error) {
	err := e.db.
		Where(&models.StudentExam{
			StudentID:        exam.StudentID,
			ExamID:           exam.ExamID,
			StudentSessionID: exam.StudentSessionID,
		}).
		FirstOrCreate(&models.StudentExam{}).Error
	if err != nil {
		return nil, errors.New("can't create exam")
	}
	return exam, nil
}

func (e *StudentExamService) DeleteStudentExam(examID string) (bool, error) {
	exam := &models.StudentExam{}
	err := e.db.Where("id=?", examID).Delete(exam).Error
	if err != nil {
		return false, errors.New("can't delete exam")
	}
	return true, nil
}

func (e *StudentExamService) GetStudentExam(examID string) (*models.StudentExam, error) {
	exam := &models.StudentExam{}
	err := e.db.Model(&models.StudentExam{}).Preload(clause.Associations).First(exam, "id=?", examID).Error
	if err != nil {
		return nil, errors.New("can get Exam ID")
	}
	return exam, nil
}

func (e *StudentExamService) GetsStudentExam() (*[]models.StudentExam, error) {
	exam := &[]models.StudentExam{}
	err := e.db.Model(&[]models.StudentExam{}).Find(exam).Error
	if err != nil {
		return nil, errors.New("can't get studentexam")
	}
	return exam, nil
}

func (e *StudentExamService) UpdateStudentExam(studentExam *models.StudentExam) (bool, error) {
	err := e.db.Where("id=?", studentExam.ID).Updates(&studentExam)
	if err != nil {
		return false, errors.New("can't update student exam")
	}
	return true, nil
}
