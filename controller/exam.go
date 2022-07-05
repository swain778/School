package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateExam(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.FormValue("subjectID"))
	b, _ := strconv.Atoi(r.FormValue("classID"))
	service := service.NewExamService()
	exam, err := service.CreateExam(&models.Exam{
		SubjectID: uint(a),
		ClassID:   uint(b),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "exam created",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "exam created",
		Data:    exam,
	})
}

func DeleteExam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewExamService()
	exam, err := service.DeleteExam(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't delete exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "exam deleted",
		Data:    exam,
	})
}

func GetExam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewExamService()
	exam, err := service.GetExam(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get exam id",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "get exam",
		Data:    exam,
	})
}

func GetsExam(w http.ResponseWriter, r *http.Request) {
	service := service.NewExamService()
	exam, err := service.GetsExam()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't get exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    exam,
	})
}

func UpdateExam(w http.ResponseWriter, r *http.Request) {
	service := service.NewExamService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	b, _ := strconv.Atoi(r.FormValue("subjectID"))
	c, _ := strconv.Atoi(r.FormValue("classID"))

	exam, err := service.UpdateExam(&models.Exam{
		ID:        uint(a),
		SubjectID: uint(b),
		ClassID:   uint(c),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "can't update exam",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    exam,
	})
}
