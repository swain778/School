package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/gorilla/mux"
)

func StudentExam(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.FormValue("studentID"))
	b, _ := strconv.Atoi(r.FormValue("examID"))
	c, _ := strconv.Atoi(r.FormValue("studentsessionID"))
	service := service.NewStudentExamService()
	studentExam, err := service.SaveStudentExamInfo(&models.StudentExam{
		StudentID:        uint(a),
		ExamID:           uint(b),
		StudentSessionID: uint(c),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create student exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "student exam created",
		Data:    studentExam,
	})
}

func DeleteStudentExam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewStudentExamService()
	studentExam, err := service.DeleteStudentExam(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete student exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "student exam deleted",
		Data:    studentExam,
	})

}

func GetStudentExam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewStudentExamService()
	studentExam, err := service.GetStudentExam(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get student exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    studentExam,
	})
}

func GetsStudentExam(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentExamService()
	studentExam, err := service.GetsStudentExam()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get student exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    studentExam,
	})
}

func UpdateStudentExam(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentExamService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	b, _ := strconv.Atoi(r.FormValue("studentID"))
	c, _ := strconv.Atoi(r.FormValue("examID"))
	d, _ := strconv.Atoi(r.FormValue("studentsessionID"))

	studentExam, err := service.UpdateStudentExam(&models.StudentExam{
		ID:               uint(a),
		StudentID:        uint(b),
		ExamID:           uint(c),
		StudentSessionID: uint(d),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "CAn't update student exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    studentExam,
	})
}
