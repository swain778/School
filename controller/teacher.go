package controller

import (
	"net/http"
	"school/models"
	service "school/service"
	"strconv"

	"github.com/gorilla/mux"
)

func AddTeacher(w http.ResponseWriter, r *http.Request) {
	service := service.NewTeacher()

	if r.FormValue("firstname") == "" || r.FormValue("lastname") == "" {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "please provide required fields",
			Data:    nil,
		})
		return
	}
	teacher, err := service.CreateTeacher(&models.Teacher{
		FirstName:  r.FormValue("firstname"),
		JoiningAt:  r.FormValue("joiningat"),
		Status:     r.FormValue("status"),
		LastName:   r.FormValue("lastname"),
		Department: r.FormValue("department"),
		DOB:        r.FormValue("dob"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create teacher",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "teacher created",
		Data:    teacher,
	})

}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewTeacher()
	teacher, err := service.DeleteTeacher(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete teacher",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "teacher deleted",
		Data:    teacher,
	})
}

func GetTeacher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewTeacher()
	teacher, err := service.GetTeacher(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    teacher,
	})
}

func GetsTeacher(w http.ResponseWriter, r *http.Request) {

	teacher, err := service.NewAttendenceService().GetsAttendence()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get teacher",
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    teacher,
	})
}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	service := service.NewTeacher()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	teacher, err := service.UpdateTeacher(&models.Teacher{
		ID:         uint(a),
		FirstName:  r.FormValue("firstname"),
		LastName:   r.FormValue("lastname"),
		Department: r.FormValue("department"),
		DOB:        r.FormValue("dob"),
		JoiningAt:  r.FormValue("joiningat"),
		Status:     r.FormValue("status"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't update teacher",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    teacher,
	})
}
