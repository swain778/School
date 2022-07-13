package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
)

type HomeworkService struct {
	db gorm.DB
}

func NewHomeworkService() *HomeworkService {
	homeworkService := new(HomeworkService)
	homeworkService.db = *config.GetDB()
	return homeworkService
}

func (h *HomeworkService) CreateHomework(homework *models.Homework) (*models.Homework, error) {

	err := h.db.Where(models.Homework{
		SubjectID:      homework.SubjectID,
		StudentID:      homework.StudentID,
		TeacherID:      homework.TeacherID,
		Class:          homework.Class,
		Homework:       homework.Homework,
		SubmissionDate: homework.SubmissionDate,
	}).FirstOrCreate(&homework).Error
	if err != nil {
		return nil, errors.New("can't create homework")
	}
	return homework, nil
}

func (h *HomeworkService) DeleteHomework(homeworkID string) (bool, error) {
	err := h.db.Where("id=?", homeworkID).Delete(&models.Homework{}).Error
	if err != nil {
		return false, errors.New("can't delete homework")
	}
	return true, nil
}
