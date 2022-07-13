package service

import (
	"errors"
	"school/config"
	"school/models"
	"school/pkg"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserService struct {
	db gorm.DB
}

func NewUserService() *UserService {
	UserService := new(UserService)
	UserService.db = *config.GetDB()
	return UserService
}

func (l *UserService) CreateUser(user *models.User) (*models.User, error) {
	err := l.db.Where(models.User{
		UserName: user.UserName,
		Password: user.Password,
		UserType: user.UserType,
	}).FirstOrCreate(&user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (l *UserService) DeleteUser(userID string) (bool, error) {
	err := l.db.Where("id=?", userID).Delete(&models.User{}).Error
	if err != nil {
		return false, errors.New("can't delete User")
	}
	return true, nil
}

func (l *UserService) GetUser(UserID string) (*models.User, error) {
	User := &models.User{}
	err := l.db.Model(&models.User{}).Preload(clause.Associations).First(User, "id=?", UserID).Error
	if err != nil {
		return nil, errors.New("can't get User")
	}
	return User, nil
}

func (l *UserService) UpdateUser(User *models.User) (bool, error) {
	err := l.db.Where("id=?", User.ID).Updates(&User).Error
	if err != nil {
		return false, errors.New("can't update User details")
	}
	return true, nil
}

func (l *UserService) UserLogin(username string, password string) (*models.Login, error) {
	userLogin := &models.Login{}
	err := l.db.Where("user_name=? AND Password=?", username, password).Preload("Role").First(userLogin).Error
	if err != nil {
		return nil, errors.New("invalid user")
	}
	return userLogin, nil
}

func (l *UserService) GetUserByModel(user *models.User) (*models.User, error) {
	userData := models.User{}
	err := l.db.Where(&user).First(&userData).Error
	if err != nil {
		return nil, errors.New("can't get password")
	}
	return &userData, nil
}

func (l *UserService) GetResetKey(resetKey string) (*models.ResetPassword, error) {
	resetPassword := models.ResetPassword{}
	err := l.db.Where(&models.ResetPassword{
		ResetKey: resetKey,
		IsUsed:   false,
	}).First(&resetPassword).Error
	if err != nil {
		return nil, err
	}
	return &resetPassword, nil
}

func (l *UserService) UpdateResetKey(resetPassword *models.ResetPassword) (bool, error) {
	err := l.db.Where(resetPassword).Updates(&models.ResetPassword{

		IsUsed: true,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (l *UserService) Auth(token string) (*models.Login, error) {
	user := models.Login{}
	if err := l.db.Where("token=?", token).
		Preload("Role").
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (l *UserService) GenerateToken(loginID uint) (string, error) {
	token := pkg.RandomString()
	err := l.db.Debug().Where("id=?", loginID).Updates(&models.Login{Token: token}).Error
	if err != nil {
		return "", err
	}
	return token, nil
}
