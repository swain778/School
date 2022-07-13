package config

import (
	"errors"
	"fmt"
	"school/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=school port=5434"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		_ = errors.New("can't connect to database")
		return nil
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(
		&models.Class{},
		&models.Complain{},
		&models.Exam{},
		&models.Fees{},
		&models.StudentAttendence{},
		&models.StudentSession{},
		&models.Student{},
		&models.Subject{},
		&models.TeacherAttendence{},
		&models.Teacher{},
		&models.ClassTeacher{},
		&models.StudentClass{},
		&models.StudentExam{},
		&models.User{},
		&models.Role{},
		&models.ResetPassword{},
		&models.User{},
		&models.Login{},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
}
