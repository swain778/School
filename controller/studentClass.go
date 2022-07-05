package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/gorilla/mux"
)

func StudentClass(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.FormValue("studentID"))
	b, _ := strconv.Atoi(r.FormValue("classID"))
	c, _ := strconv.Atoi(r.FormValue("studentsessionID"))
	service := service.NewStudentClassService()
	studentClass, err := service.CreateStudentClassService(&models.StudentClass{
		StudentID:        uint(a),
		ClassID:          uint(b),
		StudentSessionID: uint(c),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create student class",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "student class created",
		Data:    studentClass,
	})

}

func DeleteStudentClass(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewStudentClassService()
	studentClass, err := service.DeleteStudentClassService(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete student class",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "student class deleted",
		Data:    studentClass,
	})
}

func GetStudentClass(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewStudentClassService()
	studentClass, err := service.GetStudentClassService(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get student class",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "can't get student class",
		Data:    studentClass,
	})
}

func GetsStudentClass(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentClassService()
	studentClass, err := service.GetsStudentClassService()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get student class",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    studentClass,
	})

}

func UpdateStudentClass(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentClassService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	b, _ := strconv.Atoi(r.FormValue("studentID"))
	c, _ := strconv.Atoi(r.FormValue("classID"))
	d, _ := strconv.Atoi(r.FormValue("studentsessionID"))

	studentClass, err := service.UpdateStudentClassService(&models.StudentClass{
		ID:               uint(a),
		StudentID:        uint(b),
		ClassID:          uint(c),
		StudentSessionID: uint(d),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't update student class",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    studentClass,
	})
}
