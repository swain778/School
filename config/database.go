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
		&models.SessionYear{},
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
		&models.Staff{},
		&models.Salary{},
		&models.BankDetail{},
		&models.StudentHomework{},
		&models.Guardian{},
		&models.Homework{},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// func DummyData(db *gorm.DB) {

// 	db.Create(&[]models.Class{
// 		{
// 			ID:    22,
// 			Class: "5",
// 			Subject: []models.Subject{
// 				{
// 					ID:      50,
// 					Subject: "Hindi",
// 					ClassID: 22,
// 				},
// 			},
// 		},
// 		{
// 			ID:    23,
// 			Class: "5",
// 			Subject: []models.Subject{
// 				{
// 					ID:      51,
// 					Subject: "Hindi",
// 					ClassID: 23,
// 				},
// 			},
// 		},
// 	})

// }
