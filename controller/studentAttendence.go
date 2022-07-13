package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nleeper/goment"
)

func CreateAttendence(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.FormValue("studentID"))

	studate, err := goment.New(r.FormValue("date"), "YYYY-MM-DD")
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "invalid date",
			Data:    err.Error(),
		})

	}

	service := service.NewAttendenceService()
	attendence, err := service.CreateAttendence(&models.StudentAttendence{
		Attendence: r.FormValue("attendence"),
		StudentID:  uint(a),
		Date:       studate.ToTime(),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't create attendence",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "attendence created",
		Data:    attendence,
	})
}

func DeleteStudentAttendence(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewAttendenceService()
	attendence, err := service.DeleteAttendence(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete student attendence",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "student deleted",
		Data:    attendence,
	})
}

func GetStudentAttendence(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewAttendenceService()
	studentAttendence, err := service.GetAttendence(params["id"])
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
		Message: "success",
		Data:    studentAttendence,
	})
}

func GetsStudentAttendence(w http.ResponseWriter, r *http.Request) {
	service := service.NewAttendenceService()
	attendence, err := service.GetsAttendence()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get student attendence",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    attendence,
	})
}

func UpdateStudentAttendence(w http.ResponseWriter, r *http.Request) {
	service := service.NewAttendenceService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	b, _ := strconv.Atoi(r.FormValue("studentID"))

	attendence, err := service.UpdateAttendence(&models.StudentAttendence{
		Id:         uint(a),
		Attendence: r.FormValue("attendence"),
		StudentID:  uint(b),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't update student attendence",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    attendence,
	})
}
