package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StudentService struct {
	db *gorm.DB
}

func NewStudentService() *StudentService {
	studentservice := new(StudentService)
	studentservice.db = config.GetDB()
	return studentservice
}

func (s *StudentService) CreateStudent(student *models.Student) (*models.Student, error) {
	err := s.db.Where(models.Student{
		Name:        student.Name,
		DOB:         student.DOB,
		Father_Name: student.Father_Name,
		Mother_Name: student.Mother_Name,
		Address:     student.Address,
		Status:      student.Status,
	}).FirstOrCreate(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentService) DeleteStudent(studentID string) (bool, error) {
	err := s.db.Where("id = ?", studentID).Delete(&models.Student{}).Error
	if err != nil {
		return false, errors.New("can't delete student")
	}
	return true, nil
}

func (s *StudentService) GetStudent(studentID string) (*models.Student, error) {
	student := &models.Student{}
	err := s.db.
		Model(&models.Student{}).
		Preload(clause.Associations).
		First(student, "id=?", studentID).Error

	if err != nil {
		return nil, errors.New("can't get student")
	}
	return student, nil
}

func (s *StudentService) GetsStudent() (*[]models.Student, error) {
	students := &[]models.Student{}
	err := s.db.Model(&[]models.Student{}).Find(students).Error
	if err != nil {
		return nil, errors.New("can't get students")
	}
	return students, nil
}

func (s *StudentService) UpdateStudent(studentService *models.Student) (bool, error) {
	err := s.db.Where("id=?", studentService.ID).Updates(&studentService).Error
	if err != nil {
		return false, errors.New("can't update student service")
	}
	return true, nil
}
