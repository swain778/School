package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/gorilla/mux"
)

func TeacherAttendence(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.FormValue("teacherID"))
	service := service.NewTeacherAttendence()
	teacherAttendence, err := service.CreateTeacherAttendence(&models.TeacherAttendence{
		TeacherAttendence: r.FormValue("teacherattendence"),
		TeacherID:         uint(a),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create teacher attendence",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "teacher attendence created",
		Data:    teacherAttendence,
	})
}

func DeleteTeacherAttendence(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewTeacherAttendence()
	teacherAttendence, err := service.DeleteTeacherAttendence(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete teacher attendence",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "teacher attendence deleted",
		Data:    teacherAttendence,
	})
}

func GetTeacherAttendence(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewTeacherAttendence()
	teacherAttendence, err := service.GetTeacherAttendence(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "get student id",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    teacherAttendence,
	})
}

func GetsTeacherAttendence(w http.ResponseWriter, r *http.Request) {
	teacherAttendence, err := service.NewAttendenceService().GetsAttendence()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get teacher attendence",
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    teacherAttendence,
	})
}

func UpdateTeacherAttendence(w http.ResponseWriter, r *http.Request) {
	service := service.NewTeacherAttendence()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	b, _ := strconv.Atoi(r.FormValue("teacherID"))

	teacherAttendence, err := service.UpdateTeacherAttendence(&models.TeacherAttendence{
		ID:                uint(a),
		TeacherAttendence: r.FormValue("teacherattendence"),
		TeacherID:         uint(b),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "Can't update teacher attendence",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    teacherAttendence,
	})
}
