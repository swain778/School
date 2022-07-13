package service

import (
	"errors"
	"school/config"
	"school/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Roleervice struct {
	db gorm.DB
}

func NewRoleervice() *Roleervice {
	Roleervice := new(Roleervice)
	Roleervice.db = *config.GetDB()
	return Roleervice
}

func (r *Roleervice) CreateRole(role *models.Role) (*models.Role, error) {
	err := r.db.Where(models.Role{
		Role: role.Role,
	}).FirstOrCreate(&role).Error
	if err != nil {
		return nil, errors.New("can't create role")
	}
	return role, nil
}

func (r *Roleervice) DeleteRole(roleID string) (bool, error) {
	err := r.db.Where("id=?", roleID).Delete(&models.Role{}).Error
	if err != nil {
		return false, errors.New("can't delete role")
	}
	return true, nil
}

func (r *Roleervice) GetRole(roleID string) (*models.Role, error) {
	role := &models.Role{}
	err := r.db.Model(&models.Role{}).Preload(clause.Associations).First(role, "id=?", roleID)
	if err != nil {
		return nil, errors.New("can't get Role")
	}
	return role, nil
}

func (r *Roleervice) UpdateRole(role *models.Role) (bool, error) {
	err := r.db.Where("id=?", role.ID).Updates(&models.Role{}).Error
	if err != nil {
		return false, errors.New("can't update Role")
	}
	return true, nil
}
