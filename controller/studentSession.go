package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateStudentSession(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.FormValue("studentID"))
	b, _ := strconv.Atoi(r.FormValue("classID"))
	service := service.NewStudentSession()
	studentSession, err := service.CreateStudentSession(&models.StudentSession{
		Section:     r.FormValue("section"),
		SessionYear: r.FormValue("sessionyear"),
		Semister:    r.FormValue("semister"),
		Incharge:    r.FormValue("incharge"),
		StudentID:   uint(a),
		ClassID:     uint(b),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create student session",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "student session created",
		Data:    studentSession,
	})
}

func DeleteStudentSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewStudentSession()
	studentSession, err := service.DeleteStudentSession(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete student session",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "student session deleted",
		Data:    studentSession,
	})
}

func GetStudentSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewStudentSession()
	studentSession, err := service.GetStudentSession(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get student",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "get student session",
		Data:    studentSession,
	})
}

func GetsStudentSession(w http.ResponseWriter, r *http.Request) {
	//service := service.NewStudentSession()
	studentSession, err := service.NewStudentSession().GetsStudentSession()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get student session",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    studentSession,
	})
}

func UpdateStudentSession(w http.ResponseWriter, r *http.Request) {
	service := service.NewStudentSession()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	b, _ := strconv.Atoi(r.FormValue("studentID"))
	c, _ := strconv.Atoi(r.FormValue("classID"))

	studentSession, err := service.UpdateStudentSession(&models.StudentSession{
		ID:          uint(a),
		StudentID:   uint(b),
		ClassID:     uint(c),
		Section:     r.FormValue("section"),
		SessionYear: r.FormValue("sessionyear"),
		Semister:    r.FormValue("semister"),
		Incharge:    r.FormValue("incharge"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't update student session",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    studentSession,
	})
}
