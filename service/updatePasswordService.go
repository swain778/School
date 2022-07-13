package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
)

type UpdatePasswordService struct {
	db *gorm.DB
}

func NewUpdatePasswordService() *UpdatePasswordService {
	updatepasswordservice := new(UpdatePasswordService)
	updatepasswordservice.db = config.GetDB()
	return updatepasswordservice
}

func (u *UpdatePasswordService) CreateUpdatePassword(updatePassword *models.UpdatePassword) (*models.UpdatePassword, error) {
	err := u.db.Where(models.UpdatePassword{
		UserID:          updatePassword.UserID,
		Password:        updatePassword.Password,
		ConformPassword: updatePassword.ConformPassword,
	}).Create(&updatePassword).Error
	if err != nil {
		return nil, errors.New("can't create updated password")
	}
	return updatePassword, nil
}

func (u *UpdatePasswordService) NewPassword(updatePassword *models.UpdatePassword) (bool, error) {
	err := u.db.Where("id=?", updatePassword.UserID).Updates(&updatePassword).Error
	if err != nil {
		return false, errors.New("can't update password")
	}
	return true, nil
}

func (u *UpdatePasswordService) GetRandomKey(resetKey *models.ResetPassword) (*models.ResetPassword, error) {
	err := u.db.Where("id=?", resetKey.ID).First(&resetKey).Error
	if err != nil {
		return nil, errors.New("can't get reset key")
	}
	return resetKey, nil
}
