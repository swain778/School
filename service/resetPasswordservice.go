package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
)

type ResetPasswordService struct {
	db gorm.DB
}

func NewResetPasswordService() *ResetPasswordService {
	resetPasswordService := new(ResetPasswordService)
	resetPasswordService.db = *config.GetDB()
	return resetPasswordService
}

func (r *ResetPasswordService) SaveRestKey(resetPassword *models.ResetPassword) (*models.ResetPassword, error) {
	err := r.db.Create(&resetPassword).Error

	if err != nil {
		return nil, errors.New("can't reset password")
	}
	return resetPassword, nil
}

func (r *ResetPasswordService) GetRestKey(resetKey *models.ResetPassword) (*models.ResetPassword, error) {
	err := r.db.Where("id=?", resetKey.ID).First(&resetKey).Error
	if err != nil {
		return nil, errors.New("can't get reset key")
	}
	return resetKey, nil
}
