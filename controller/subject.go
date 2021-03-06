package controller

import (
	"net/http"
	"school/models"
	"school/service"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
)

func AddSubject(w http.ResponseWriter, r *http.Request) {
	service := service.NewSubjectService()
	subject, err := service.CreateSubject(&models.Subject{
		Subject: r.FormValue("subject"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    subject,
	})
}

func GetsSubject(w http.ResponseWriter, r *http.Request) {
	service := service.NewSubjectService()
	subject, err := service.GetsSubject()
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Success",
		Data:    subject,
	})
}

func GetSubject(w http.ResponseWriter, r *http.Request) {
	service := service.NewSubjectService()
	params := mux.Vars(r)
	subject, err := service.GetSubject(params["id"])
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    err.Error(),
		})
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Success",
		Data:    subject,
	})
}

func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	service := service.NewSubjectService()
	subject, err := service.DeleteSubject(chi.URLParam(r, "id"))
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    err.Error(),
		})
		return
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "success",
		Data:    subject,
	})
}

func UpdateSubject(w http.ResponseWriter, r *http.Request) {
	service := service.NewSubjectService()
	params := mux.Vars(r)
	a, _ := strconv.Atoi(params["id"])
	subject, err := service.UpdateSubject(&models.Subject{
		ID:      uint(a),
		Subject: r.FormValue("subject"),
	})
	if err != nil {
		ApiResponse(w, &Res{
			Code:    900,
			Message: "error",
			Data:    err.Error(),
		})
	}
	ApiResponse(w, &Res{
		Code:    200,
		Message: "Success",
		Data:    subject,
	})
}
